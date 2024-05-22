import React, { useState } from 'react';
import { RegisterStudentAPICall } from '../../../action/Auth/register';
import { LoginStudentAPICall } from '../../../action/Auth/login';


function RegisterStudent(props) {
    const [formData, setFormData] = useState({
        login: '',
        password: '',
        username: '',
        system_password: ''
    });
    const [showAdditionalFields, setShowAdditionalFields] = useState(false);

    const handleChange = (e) => {
        const { name, value } = e.target;
    
        if (name === 'login') {
            const numericValue = value.replace(/\D/g, ''); 
            if (numericValue.length > 12) return; 
            setFormData({
                ...formData,
                [name]: numericValue
            });
        } else {
            setFormData({
                ...formData,
                [name]: value
            });
        }
    };
    
    

    const next = () => {
        if (formData.login.length != 12 || formData.password.length == 0){
            return 
        }
        setShowAdditionalFields(true)
    }

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            let regResp = await RegisterStudentAPICall(formData);
            props.setCode(regResp.code);
            props.setMessage(regResp.message);
    
            if (regResp.code === 200) {
                let regResp = await LoginStudentAPICall(formData);
                props.setCode(regResp.code);
                props.setMessage(regResp.message);
        
                if (regResp.code === 200) {
                    localStorage.setItem('token', regResp.token);
                    window.location.href = '/home';
                      }
            }
        } catch (error) {
            props.setCode(400);
            props.setMessage("Login failed: " + error.message);
        }
    };   

    async function register() {
        let regResp = RegisterStudentAPICall(formData)
        props.setCode((await regResp).code)
        props.setMessage((await regResp).message)
        
    }    

    return (
        <>
            <form className="card-body" onSubmit={handleSubmit}>
                {!showAdditionalFields ? (
                    <>
                        <div className="form-control">
                            <label className="label">
                                <span className="label-text">Login (12 digits)</span>
                            </label>
                            <input
                                type="text"
                                name="login"
                                placeholder="sms's login"
                                className="input input-bordered"
                                value={formData.login}
                                onChange={handleChange}
                                required
                            />
                        </div>
                        <div className="form-control">
                            <label className="label">
                                <span className="label-text">Password</span>
                            </label>
                            <input
                                type="password"
                                name="password"
                                placeholder="sms's password"
                                className="input input-bordered"
                                value={formData.password}
                                onChange={handleChange}
                                required
                            />
                        </div>
                        <button className="btn btn-primary mt-3" type="button" onClick={next}> Next</button>
                    </>
                ) : (
                    <>
                        <div className="form-control">
                            <label className="label">
                                <span className="label-text">Username</span>
                            </label>
                            <input
                                type="text"
                                name="username"
                                placeholder="lms's username"
                                className="input input-bordered"
                                                                
                                value={formData.username}
                                onChange={handleChange}
                                required
                            />
                        </div>
                        <div className="form-control">
                            <label className="label">
                                <span className="label-text">System Password</span>
                            </label>
                            <input
                                type="password"
                                name="system_password"
                                placeholder="lms's password"
                                className="input input-bordered"
                                value={formData.systemPassword}
                                onChange={handleChange}
                                required
                            />
                        </div>
                        <button className="btn btn-primary mt-3" type="button" onClick={register}>Submit</button>
                    </>
                )}
                <div className="justify-center align-center flex mt-2">
                    <a href="#" onClick={(e) => { e.preventDefault(); props.set(null); }}>Go Back</a>
                </div>
            </form>
        </>
    );
}

export default RegisterStudent;
