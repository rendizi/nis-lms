import React from "react";
import { isTokenValid } from "../../action/Auth/validate";
import { LoginStudentAPICall } from "../../action/Auth/login";
import { LoginTeacherAPICall } from "../../action/Auth/login";
import Navbar from "../Components/Navbar";

function Home(){
    const [token, setToken] = React.useState(null)

    React.useEffect(() => {
        const fetchData = async () => {
            let storedToken = localStorage.getItem("token");
            if (storedToken === null) {
                window.location.href = "/";
            } else {
                if (isTokenValid(storedToken)) {
                    setToken(storedToken);
                } else {
                    let storedLogin = localStorage.getItem("login");
                    let storedPassword = localStorage.getItem("password");
                    let role = localStorage.getItem("role");
                    if (storedLogin === null || storedPassword === null || role === null){
                        console.log("no login data")
                        localStorage.setItem("token", null);
                        window.location.href = "/";
                    } else {
                        let logResp;
                        if (role === "student"){
                            logResp = await LoginStudentAPICall({ login: storedLogin, password: storedPassword });                        
                        }else{
                            logResp = await LoginTeacherAPICall({ login: storedLogin, password: storedPassword });    
                        }
                        if (logResp.code == 200){
                            localStorage.setItem("token", logResp.token)
                        }else{
                            window.location.href = "/home"
                        }
                    }
                }
            }
        };
    
        fetchData();
    }, []);
    

    return (
        <>
            <Navbar username={localStorage.getItem("login")}/>
        </>
    )
}

export default Home 