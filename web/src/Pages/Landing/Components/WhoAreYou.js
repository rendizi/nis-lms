import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUserGraduate, faChalkboardTeacher } from '@fortawesome/free-solid-svg-icons';

function WhoAreYou(props){
    return (
        <>
        <div className="flex justify-center items-center">
                    <h2 className="text-2xl font-bold">Who are you, warrior?</h2>
                </div>

                <div className="flex justify-around mt-4">
                    <div className="text-center w-40">
                        <button 
                            type="button" 
                            className="flex flex-col items-center border-2 border-transparent hover:border-blue-500 p-4 rounded-lg w-full"
                            onClick={() => props.setRole("teacher")}
                        >
                            <FontAwesomeIcon icon={faChalkboardTeacher} size="6x" className="text-blue-500" />
                            <span className="mt-2 text-lg">Teacher</span>
                        </button>
                    </div>
                    <div className="text-center w-40">
                        <button 
                            type="button" 
                            className="flex flex-col items-center border-2 border-transparent hover:border-red-500 p-4 rounded-lg w-full"
                            onClick={() => props.setRole("student")}
                        >
                            <FontAwesomeIcon icon={faUserGraduate} size="6x" className="text-red-500" />
                            <span className="mt-2 text-lg">Student</span>
                        </button>
                    </div>
                </div>
                <div className="justify-center align-center flex mt-4">
                    <a href="#" onClick={(e) => { e.preventDefault(); props.set(true); }}>
                        Already have an account?
                    </a>
                </div>
            </>
    )
}

export default WhoAreYou