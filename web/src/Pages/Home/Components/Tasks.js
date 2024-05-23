import React, { useState, useEffect } from "react";
import { SearchTask } from "../../../action/Tasks/search";

function Tasks() {
    const [searchInput, setSearchInput] = useState('');
    const [difficulty, setDifficulty] = useState('');
    const [tasks, setTasks] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [currentPage, setCurrentPage] = useState(1); // Track current page

    const handleSearch = async () => {
        setLoading(true);
        setError(null);
        try {
            const result = await SearchTask({ difficulty, title: searchInput, page: currentPage }); 
            setTasks(result);
        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    useEffect(() => {
        handleSearch(); 
    }, [currentPage]);

    return (
        <div className="flex justify-center items-center h-full">
            <div className="flex flex-col justify-between items-center h-full bg-neutral   w-1/2 rounded-2xl relative">
                <div className="absolute top-0 right-0 mt-1 mr-1">
                <div className="join">
                    <div>
                        <input
                            className="input input-bordered join-item  "
                            placeholder="Search"
                            value={searchInput}
                            onChange={(e) => setSearchInput(e.target.value)}
                        />
                    </div>
                    <select
                        className="select select-bordered join-item "
                        value={difficulty}
                        onChange={(e) => setDifficulty(e.target.value)}
                    >
                        <option value="" disabled selected>Difficulty</option>
                        <option value="easy">Easy</option>
                        <option value="medium">Medium</option>
                        <option value="hard">Hard</option>
                        <option value="">All</option>
                    </select>
                    <div className="indicator">
                        <button className="btn join-item " onClick={handleSearch}>Search</button>
                    </div>
                </div>

                </div>
                <div className="flex-1 flex flex-col justify-center items-center w-full p-4 space-y-2 mt-5">
                    {loading && <p>Loading...</p>}
                    {error && <p className="text-red-500">{error}</p>}
                    {tasks && tasks.length > 0 && tasks.map((task) => (
                        <div key={task.id} className="bg-neutral-content w-full p-4 rounded-lg shadow-md flex justify-between items-center h-10 cursor-pointer" onClick={()=>{window.location.href = "/t/"+task.id}}>
                            <div className="flex items-center">
                                <h2 className="text-xl font-bold mr-2 text-gray-800">{task.id}. {task.title}</h2>
                                <sub><p className=" text-sm text-gray-600">by {task.author}</p></sub>
                            </div>
                            <p className={task.difficulty === "easy" ? "text-success " :
                             task.difficulty === "medium" ? "text-warning " : "text-error "}>{task.difficulty}</p>
                        </div>
                    ))}
                </div>

                <div className="join mb-1">
                    <button className="join-item btn" onClick={() => {
                        if (currentPage <= 1){
                            return
                        }
                        setCurrentPage(currentPage - 1)
                        }}>«</button> {/* Previous page */}
                    <button className="join-item btn">Page {currentPage}</button> {/* Display current page */}
                    <button className="join-item btn" onClick={() => setCurrentPage(currentPage + 1)}>»</button> {/* Next page */}
                </div>
            </div>
        </div>
    );
}

export default Tasks;
