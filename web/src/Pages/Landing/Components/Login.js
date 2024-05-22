import React, { useState } from "react";
import { LoginStudentAPICall,LoginTeacherAPICall } from "../../../action/Auth/login";

function Login(props){
    const [formData, setFormData] = useState({
        login: '',
        password: '',
    });
    const [userType, setUserType] = useState('student');
    const [isChecked, setIsChecked] = useState(false);

    const handleCheckboxChange = (event) => {
        setIsChecked(event.target.checked);
      };

    React.useEffect(()=>{
        async function log(){
        let storedLogin = localStorage.getItem("login")
        let storedPassword = localStorage.getItem("password")
        let storedRole = localStorage.getItem("role")

        if (storedLogin !== null && storedPassword !== null && storedRole!== null){
            let regResp;
            if (userType === 'student') {
                regResp = await LoginStudentAPICall({login: storedLogin, password: storedPassword});
            } else if (userType === 'teacher') {
                regResp = await LoginTeacherAPICall({login: storedLogin, password: storedPassword});
            }
        
            if (regResp.code === 200) {
                localStorage.setItem('token', regResp.token);
                window.location.href = '/home';
            }
        }}
        log()

    },[])

    const handleChange = (e) => {
        const { name, value } = e.target;

        setFormData({
            ...formData,
            [name]: value
        });
    };

    const handleUserTypeChange = (e) => {
        setUserType(e.target.value);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            let regResp;
            if (userType === 'student') {
                regResp = await LoginStudentAPICall(formData);
            } else if (userType === 'teacher') {
                regResp = await LoginTeacherAPICall(formData);
            }
            localStorage.setItem("role", userType)

            if (!isChecked){
                localStorage.removeItem("password")
            }
    
            props.setCode(regResp.code);
            props.setMessage(regResp.message);
    
            if (regResp.code === 200) {
                localStorage.setItem('token', regResp.token);
                window.location.href = '/home';
            }
        } catch (error) {
            props.setCode(400);
            props.setMessage("Login failed: " + error.message);
        }
    };
    

    return (
        <div className="card shrink-0 w-full max-w-sm bg-base-100">
            <form className="card-body" onSubmit={handleSubmit}>
                <div className="form-control">
                    <label className="label">
                        <span className="label-text">Login</span>
                    </label>
                    <input type="text" name="login" placeholder="login" className="input input-bordered" required onChange={handleChange} value={formData.login}/>
                </div>
                <div className="form-control">
                    <label className="label">
                        <span className="label-text">Password</span>
                    </label>
                    <input type="password" name="password" placeholder="password" className="input input-bordered" required onChange={handleChange} value={formData.password}/>
                </div>
                <div className='flex justify-center items-center mt-2'>
                    <label className="flex items-center">
                        <input type="radio" name="userType" className="radio" value="student" checked={userType === 'student'} onChange={handleUserTypeChange}/>
                        <span className="ml-1">Student</span>
                    </label>
                    <label className="flex items-center ml-5">
                        <input type="radio" name="userType" className="radio" value="teacher" checked={userType === 'teacher'} onChange={handleUserTypeChange}/>
                        <span className="ml-1">Teacher</span>
                    </label>
                </div>
                <div className="form-control flex justify-center items-center">
                    <label className="label cursor-pointer">
                        <input
                        type="checkbox"
                        className="checkbox"
                        checked={isChecked}
                        onChange={handleCheckboxChange}
                        />
                        <span className="label-text ml-2">Remember me</span>
                    </label>
                </div>
                <div className="form-control mt-2">
                    <button className="btn btn-primary" type="submit">Login</button>
                    <div className="justify-center align-center flex mt-2">
                        <a href="#" onClick={(e) => {e.preventDefault(); props.set(false);}}>Don't have an account yet?</a>
                    </div>
                </div>
            </form>
        </div>
    )
}

export default Login;
