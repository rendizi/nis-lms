import React, { useState } from 'react';

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

    const handleSubmit = (e) => {
        e.preventDefault();

        console.log(formData)
    };

    async function register() {
        console.log(formData);
        try {
            const response = await fetch('http://localhost:8080/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(formData)
            });
    
            if (response.status === 200) {
                console.log('success');
            } else {
                console.log('Request failed with status:', response.status);
            }
        } catch (error) {
            console.error('Error:', error);
        }
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
