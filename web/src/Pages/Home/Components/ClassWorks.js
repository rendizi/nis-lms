import React from "react";
import { FetchStudentsClassWork } from "../../../action/Classwork/get";

function ClassWorks() {
    const [classWork, setClassWork] = React.useState(null);
    const [currentIndex, setCurrentIndex] = React.useState(0);
    const itemsPerPage = 4;

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

    React.useEffect(() => {
        const fetchData = async () => {
            let user = localStorage.getItem("login");
            try {
                let classworkData = await FetchStudentsClassWork(user);
                if (Array.isArray(classworkData) && classworkData.length > 0) {
                    console.log('Classwork fetched successfully:', classworkData);
                    setClassWork(classworkData);
                } else {
                    console.log('No classwork found for the user.');
                }
            } catch (error) {
                console.error('Error fetching classwork:', error);
            }
        };

        fetchData();
    }, []);

    if (classWork === null) {
        return (
            <div className="flex justify-center align-center">
                <h1 className="text-center mt-2">ClassWorks</h1>
                <span className="loading loading-spinner loading-lg"></span>
            </div>
        );
    }

    return (
        <div className="flex flex-col justify-center items-center ">
            <h1 className="text-center mt-2">ClassWorks</h1>
            <div className="flex justify-between items-center w-full mt-5 w-max">
                <div className="flex justify-center align-center ml-5">
                    
                    <div className="flex flex-nowrap overflow-x-auto">
                        {displayClassWork()}
                    </div>
                    
                </div>
            </div>
        </div>
    );
}

export default ClassWorks;
