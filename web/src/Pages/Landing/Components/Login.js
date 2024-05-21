function Login(props){
    return (
        <div className="card shrink-0 w-full max-w-sm bg-base-100">
                        <form className="card-body">
                            <div className="form-control">
                                <label className="label">
                                    <span className="label-text">Login</span>
                                </label>
                                <input type="text" placeholder="login" className="input input-bordered" required />
                            </div>
                            <div className="form-control">
                                <label className="label">
                                    <span className="label-text">Password</span>
                                </label>
                                <input type="password" placeholder="password" className="input input-bordered" required />
                                </div>
                                <div className='flex justify-center items-center mt-2'>
                                    <label className="flex items-center">
                                        <input type="radio" name="radio-1" className="radio" checked />
                                        <span className="ml-1">Student</span>
                                    </label>
                                    <label className="flex items-center ml-5">
                                        <input type="radio" name="radio-1" className="radio" />
                                        <span className="ml-1">Teacher</span>
                                    </label>
                                </div>

                                <div className="form-control mt-6">
                                <button className="btn btn-primary">Login</button>
                                <div className="justify-center align-center flex mt-2">
                                <a href="#" onClick={(e) => {e.preventDefault(); props.set(false);}}>Don't have an account yet?</a>                                </div>
                                </div>
                        </form>
                    </div>
    )
}

export default Login