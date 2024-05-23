import React from "react";
import { GetStudentData } from "../../action/others/profile";
import { GetStudentSolutions } from "../../action/Tasks/get_users_solutions";
import Navbar from "../Components/Navbar";

function Profile() {
    const [username, setUsername] = React.useState(null);
    const [data, setData] = React.useState(null);
    const [solutions, setSolutions] = React.useState([]);
    const [page, setPage] = React.useState(1);

    React.useEffect(() => {
        const currentRoute = window.location.pathname;

        if (currentRoute.startsWith("/u/")) {
            const remainingPart = currentRoute.substring(3); 
            console.log("Remaining part:", remainingPart);
            setUsername(remainingPart);
        } else {
            console.log("Current route doesn't start with '/u/'");
        }
    }, []);

    React.useEffect(() => {
        if (username) {
            async function fetchData() {
                const resp = await GetStudentData(username);
                console.log(resp);
                setData(resp);
                const sols = await GetStudentSolutions(username, page);
                setSolutions(sols);
            }
            fetchData();
        }
    }, [username, page]);

    const handlePreviousPage = () => {
        if (page > 1) {
            setPage(page - 1);
        }
    };

    const handleNextPage = () => {
        setPage(page + 1);
    };

    return (
        <>
        <Navbar />
        <div className="flex flex-col items-center justify-center py-8">
            <div className="flex w-full max-w-4xl border shadow-lg">
                <div className="bg-info-content shadow-md rounded-lg p-8 w-1/2  rounded-lg" >
                    {data !== null && data.stats && (
                        <>
                            <h1 className="text-3xl font-bold text-center mb-6">{data.id}. {data.login}</h1>
                            <div>
                                <div className="mb-4">
                                    <p className="text-lg font-medium">Email: <span className="text-gray-600">{data.email}</span></p>
                                    <p className="text-lg font-medium">Klass: <span className="text-gray-600">{data.klass}</span></p>
                                    <p className="text-lg font-medium">School: <span className="text-gray-600">{data.school}</span></p>
                                </div>
                                <div>
                                    <p className="text-xl font-semibold mb-2">Stats:</p>
                                    <ul className="list-disc list-inside">
                                        <li className="text-lg">Solved: <span className="text-gray-600">{data.stats.solved}</span></li>
                                        <li className="text-lg">LeetCode: <span className="text-gray-600">{data.stats.leetcode}</span></li>
                                        <li className="text-lg">Badges: <span className="text-gray-600">{data.stats.badges}</span></li>
                                        <li className="text-lg">Rating: <span className="text-gray-600">{data.stats.rating}</span></li>
                                        <li className="text-lg">Rank: <span className="text-gray-600">{data.stats.rank}</span></li>
                                    </ul>
                                </div>
                            </div>
                        </>
                    )}
                </div>
                <div className="w-1/2 p-8 bg-info-content">
                    <h2 className="text-2xl font-bold mb-4">Solved</h2>
                    {solutions && solutions.length > 0 ? (
                        <>
                            <ul className="list-disc list-inside">
                                {solutions.map((solution, index) => (
                                    <li key={index} className="text-lg mb-2">
                                        {solution.id}.{solution.title} - {solution.time}
                                    </li>
                                ))}
                            </ul>
                            
                        </>
                    ) : (
                        <p>No solutions found.</p>
                    )}
                    <div className="flex justify-between mt-4">
                                <button 
                                    onClick={handlePreviousPage} 
                                    className="bg-blue-500 text-white px-4 py-2 rounded-md"
                                    disabled={page === 1}
                                >
                                    Previous
                                </button>
                                <button 
                                    onClick={handleNextPage} 
                                    className="bg-blue-500 text-white px-4 py-2 rounded-md"
                                >
                                    Next
                                </button>
                            </div>
                </div>
            </div>
        </div>
        </>
    );
}

export default Profile;
