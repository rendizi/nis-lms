import React, { useState, useEffect } from 'react';

import WhoAreYou from './WhoAreYou';
import RegisterStudent from './RegisterStudent';
import RegisterTeacher from './RegisterTeacher';

function Register(props) {
    const [role, setRole] = useState(null);

    useEffect(() => {
        
    }, [role]);

    return (
        <div className="card shrink-0 w-full max-w-sm bg-base-100">
            <form className="card-body">
                {role === null && <WhoAreYou set={props.set} setRole={setRole}/> || role === "student" && <RegisterStudent set={setRole}/> || <RegisterTeacher set={setRole}/>}                
            </form>
        </div>
    );
}

export default Register;
