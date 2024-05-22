import '@fortawesome/fontawesome-free/css/all.min.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import React from 'react';

function Navbar(props){
    const [profile, setProfile] = React.useState()


React.useEffect(() => {
    let role = localStorage.getItem("role")
    if (role == "student"){
       setProfile("/u/"+props.username)}
    else{
        setProfile("/t/"+props.username)
    }
    }, []);
    return (
        <div className="navbar bg-base-100">
            <div className="flex-1">
                <a className="btn btn-ghost text-xl" href="/">
                    <i className="fas fa-book"></i> LMS
                </a>
            </div>
            <div className="flex-none">
                <ul className="menu menu-horizontal px-1">
                    <li>
                        <details>
                            <summary>
                                <i className="fas fa-bars"></i> More
                            </summary>
                            <ul className="p-2 bg-base-100 rounded-t-none">
                                <li><a href={profile}><i className="fas fa-user-circle"></i> Profile</a></li>
                                <li><a href="/top"><i className="fas fa-chevron-up"></i> Top</a></li>
                                <li><a href="/home"><i className="fas fa-home"></i> Home</a></li>
                                <li><a href="#" onClick={()=>{
                                    localStorage.removeItem("token")
                                    localStorage.removeItem("login")
                                    localStorage.removeItem("password")
                                    window.location.href = "/"
                                }}><i className="fas fa-sign-out-alt"></i> Log out</a></li>
                            </ul>
                        </details>
                    </li>
                </ul>
            </div>
        </div>
    )
}

export default Navbar;
