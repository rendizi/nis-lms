import { faExclamationCircle } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';

function RegisterTeacher(props){
    return (
        <>
         <form className="card-body">
         <div className="form-control text-center">
         <h2 className="text-2xl font-bold">
        <FontAwesomeIcon icon={faExclamationCircle} className="text-red-600 mr-2" />
        Ooops....
    </h2>
    <p className="text-base">We don't have a registration form for teachers yet. Please, contact us:
        <br/>
        <span>Telegram:</span> <a href="https://t.me/rendizi" className="underline"> @rendizi</a>, 
        <br/>
        <span>Mail:</span> <a href="mailto:baglanov_a0930@akb.nis.edu.kz" className="underline">baglanov_a0930@akb.nis.edu.kz</a> 
        <br/> and we will add you manually.
    </p>
    <div className="justify-center align-center flex mt-4">
        <a href="#" className="text-blue-500 hover:underline" onClick={(e) => {e.preventDefault(); props.set(null);}}>Go Back</a> 
    </div>
</div>

                        </form>
        </>
    )
}
export default RegisterTeacher