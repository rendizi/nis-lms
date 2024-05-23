import React from "react";
import { FetchStudentsClassWork } from "../../../action/Classwork/get";
import { FetchTeachersClassWork } from "../../../action/Classwork/get";

function ClassWorks() {
    const [classWork, setClassWork] = React.useState(undefined);
    const [currentIndex, setCurrentIndex] = React.useState(0);
    const itemsPerPage = 4;
    const [allDone, setAllDone] = React.useState(null)

    const displayClassWork = () => {
        return (
            <div className="flex items-center">
                {classWork.length > itemsPerPage && <button className="join-item btn mr-4" onClick={handlePrevClick}>«</button>}
                <div className="flex overflow-x-auto">
                    {classWork.slice(currentIndex, currentIndex + itemsPerPage).map((work, index) => (
                        <div key={index} className="card bg-neutral  rounded-box p-4 mr-4 w-64">
                            <p className="font-bold">{work.title}</p>
                            <p className="text-sm mt-2">{work.description}</p>
                            <p className="text-sm">Deadline: {work.deadline}</p>
                            <p className="text-sm">Tasks: {work.tasks.join(', ')}</p>
                        </div>
                    ))}
                </div>
                {classWork.length > itemsPerPage && <button className="join-item btn ml-4" onClick={handleNextClick}>»</button>}
            </div>
        )
        
    };

    const handlePrevClick = () => {
        if (currentIndex > 0) {
            setCurrentIndex(currentIndex - itemsPerPage);
        }
    };

    const handleNextClick = () => {
        if (classWork && currentIndex + itemsPerPage < classWork.length) {
            setCurrentIndex(currentIndex + itemsPerPage);
        }
    };

    const addNew = () => {
        document.getElementById('my_modal_1').showModal()
    }

    const [formData, setFormData] = React.useState({
        title: '',
        description: '',
        deadline: '',
        students: [],
        tasks: []
      });
    
      const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData(prevState => ({
          ...prevState,
          [name]: value
        }));
      };
    
      const handleSubmit = (e) => {
        e.preventDefault();
      
        const token = localStorage.getItem("token");
      
        if (token) {
          const transformedData = {
            "Tasks": formData.tasks.split(",").map(Number),
            "Deadline": formData.deadline,
            "Title": formData.title,
            "Description": formData.description,
            "Students": formData.students.split(",").map(Number)
          };
      
          fetch("http://localhost:8080/classwork", {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
              "Authorization": `${token}`
            },
            body: JSON.stringify(transformedData)
          })
          .then(response => {
            if (response.ok) {
              return response.json();
            } else {
              throw new Error("Failed to add classwork");
            }
          })
          .then(data => {
            console.log(data);
          })
          .catch(error => {
            console.error(error.message);
          });
        } else {
          console.error("Token not found in localStorage");
        }
      };
      
      

    React.useEffect(() => {
        const fetchData = async () => {
            let user = localStorage.getItem("login");
            try {
                let classworkData
                if (localStorage.getItem("role") === "teacher"){
                    classworkData = await FetchTeachersClassWork(user)
                }else{
                classworkData = await FetchStudentsClassWork(user);}
                if (Array.isArray(classworkData) && classworkData.length > 0) {
                    console.log('Classwork fetched successfully:', classworkData);
                    setClassWork(classworkData);
                } else {
                    setAllDone(true)
                }
            } catch (error) {
                console.error('Error fetching classwork:', error);
            }
        };

        fetchData();
    }, []);

    if (allDone !== null && allDone){
        return (
            <>
            <div className="flex flex-col justify-center items-center ">
            <h1 className="text-center mt-2">ClassWorks</h1> 
                <div className="flex justify-between items-center w-full mt-5 w-max">
                <div className="flex justify-center align-center ml-5">
                    
                    <div className="flex flex-nowrap overflow-x-auto">
                    <div className="flex items-center">
                
                        <div className="flex overflow-x-auto">
                        
                                <div className="card bg-neutral  rounded-box p-4 mr-4 w-64">
                                    <p className="font-bold">There will be shown title</p>
                                    <p className="text-sm mt-2">Description</p>
                                    <p className="text-sm">Deadline</p>
                                    <p className="text-sm">Tasks</p>
                                </div>
                        </div>
                    </div>
                    </div>
                    
                </div>
                </div>
        </div>
        <dialog id="my_modal_1" className="modal">
        <div className="modal-box">
    <h3 className="font-bold text-lg">Hello!</h3>
    <p className="py-4">Press ESC key or click outside to close</p>
  </div>
  <form method="dialog" className="modal-backdrop">
    <button>close</button>
  </form>
</dialog>
        </>
        )
    }

    if (classWork === undefined) {
        return (
            <div className="flex justify-center align-center">
                <h1 className="text-center mt-2">ClassWorks</h1>
                <span className="loading loading-spinner loading-lg"></span>
            </div>
        );
    }

    return (
        <>
        <div className="flex flex-col justify-center items-center ">
            <h1 className="text-center mt-2">ClassWorks</h1>
            {localStorage.getItem("role") === "teacher" && <a href="#" onClick={addNew}>create new</a>}
            <div className="flex justify-between items-center w-full mt-5 w-max">
                <div className="flex justify-center align-center ml-5">
                    
                    <div className="flex flex-nowrap overflow-x-auto">
                        {displayClassWork()}
                    </div>
                    
                </div>
            </div>
        </div>
        <dialog id="my_modal_1" className="modal">
        <div className="modal-box">
        <form onSubmit={handleSubmit} className="w-full max-w-md">
  <div className="mb-4">
    <label className="block text-white text-sm font-bold mb-2" htmlFor="title">
      Title
    </label>
    <input
      className="shadow appearance-none border rounded w-full py-2 px-3 text-white leading-tight focus:outline-none focus:shadow-outline"
      id="title"
      type="text"
      name="title"
      placeholder="Enter title"
      value={formData.title}
      onChange={handleChange}
    />
  </div>
  <div className="mb-4">
    <label className="block text-white text-sm font-bold mb-2" htmlFor="description">
      Description
    </label>
    <textarea
      className="shadow appearance-none border rounded w-full py-2 px-3 text-white leading-tight focus:outline-none focus:shadow-outline"
      id="description"
      name="description"
      placeholder="Enter description"
      value={formData.description}
      onChange={handleChange}
    />
  </div>
  <div className="mb-4">
    <label className="block text-white text-sm font-bold mb-2" htmlFor="deadline">
      Deadline
    </label>
    <input
      className="shadow appearance-none border rounded w-full py-2 px-3 text-white leading-tight focus:outline-none focus:shadow-outline"
      id="deadline"
      type="date"
      name="deadline"
      value={formData.deadline}
      onChange={handleChange}
    />
  </div>
  {/* Input for students */}
  <div className="mb-4">
    <label className="block text-white text-sm font-bold mb-2" htmlFor="students">
      Students (comma-separated list of numbers)
    </label>
    <input
      className="shadow appearance-none border rounded w-full py-2 px-3 text-white leading-tight focus:outline-none focus:shadow-outline"
      id="students"
      type="text"
      name="students"
      placeholder="Enter students"
      value={formData.students}
      onChange={handleChange}
    />
  </div>
  {/* Input for tasks */}
  <div className="mb-4">
    <label className="block text-white text-sm font-bold mb-2" htmlFor="tasks">
      Tasks (comma-separated list of numbers)
    </label>
    <input
      className="shadow appearance-none border rounded w-full py-2 px-3 text-white leading-tight focus:outline-none focus:shadow-outline"
      id="tasks"
      type="text"
      name="tasks"
      placeholder="Enter tasks"
      value={formData.tasks}
      onChange={handleChange}
    />
  </div>
  <div className="flex items-center justify-between">
    <button
      className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
      type="submit"
    >
      Submit
    </button>
  </div>
</form>

    </div>
        <form method="dialog" className="modal-backdrop">
            <button>close</button>
        </form>
        </dialog>
        </>
    );
}

export default ClassWorks;
