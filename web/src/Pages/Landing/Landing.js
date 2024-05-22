import { FaGithub } from 'react-icons/fa';
import Login from './Components/Login';
import React from 'react';
import Register from './Components/Register';

function Landing() {
    const [haveAnAccount, setHaveAnAccount] = React.useState(false)
    const [code, setCode] = React.useState(null)
    const [message, setMessage] = React.useState(null)


    return (
        <>
        <div className="hero min-h-screen bg-base-200">
            <div className="hero-content flex-col lg:flex-row">
                <img src="hero.png" className="max-w-sm rounded-lg" alt="Hero" />
                <div>
                    <h1 className="text-5xl font-bold">Learning Management System</h1>
                    <p className="py-6">Solve programming problems, do homework and compile your code there! Project is open-source btw, you can contribute if you want to</p>
                    <div className="flex items-center">
                        <button className="btn btn-primary mr-5" onClick={() => document.getElementById('my_modal_2').showModal()}>Get started!</button>
                        <a href="https://github.com/rendizi/nis-lms" target="_blank" rel="noopener noreferrer" className="btn btn-secondary flex items-center">
                            <FaGithub size={24} className="mr-1" />
                            GitHub
                        </a>
                    </div>
                </div>
            </div>
            <dialog id="my_modal_2" className="modal">
                <div className="modal-box flex justify-center items-center">
                    {haveAnAccount ? <Login set={setHaveAnAccount} setCode={setCode} setMessage={setMessage} /> : <Register set={setHaveAnAccount} setCode={setCode} setMessage={setMessage}/>}
                </div>
    
                <form method="dialog" className="modal-backdrop">
                    <button>close</button>
                </form>
            </dialog>
        </div>
        {code !== null && code !== 200 ?
        <div role="alert" className="alert alert-error fixed bottom-0 right-0 mb-4 mr-4 md:w-1/2 w-full z-50">
            <svg xmlns="http://www.w3.org/2000/svg" className="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
            <span>Error! {message === null ? "Something went wrong. Try again later" : message}</span>
        </div> : code == 200 &&
         <div role="alert" className="alert alert-success fixed bottom-0 right-0 mb-4 mr-4 md:w-1/2 w-full z-50">
         <svg xmlns="http://www.w3.org/2000/svg" className="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24">
             <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" />
         </svg>
         <span>Success! You will be redirected soon</span>
     </div>
}
    </>
    
    );
}

export default Landing;
