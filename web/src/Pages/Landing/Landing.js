import { FaGithub } from 'react-icons/fa';
import Login from './Components/Login';
import React from 'react';
import Register from './Components/Register';

function Landing() {
    const [haveAnAccount, setHaveAnAccount] = React.useState(false)



    return (
        <div className="hero min-h-screen bg-base-200">
            <div className="hero-content flex-col lg:flex-row">
                <img src="hero.png" className="max-w-sm rounded-lg" alt="Hero" />
                <div>
                    <h1 className="text-5xl font-bold">Learning Management System</h1>
                    <p className="py-6">Solve programming problems, do homework and compile your code there! Project is open-source btw, you can contribute if you want to</p>
                    <div className="flex items-center">
                        <button className="btn btn-primary mr-5" onClick={()=>document.getElementById('my_modal_2').showModal()}>Get started!</button>
                        <a href="https://github.com/rendizi/nis-lms" target="_blank" rel="noopener noreferrer" className="btn btn-secondary flex items-center">
                            <FaGithub size={24} className="mr-1" />
                            GitHub
                        </a>
                    </div>
                </div>
            </div>
            <dialog id="my_modal_2" className="modal">
                <div className="modal-box flex justify-center items-center">
                    {haveAnAccount ? <Login set={setHaveAnAccount}/> : <Register set={setHaveAnAccount}/>}
                </div>

                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </div>
    );
}

export default Landing;
