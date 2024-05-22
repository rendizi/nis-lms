import React from "react";
import { GetStudentData } from "../../action/others/profile";

function Profile(){
    const [username, setUsername] = React.useState(null)
    const [data, setData] = React.useState(null)

    React.useEffect(()=>{
        const currentRoute = window.location.pathname;

        if (currentRoute.startsWith("/u/")) {
            const remainingPart = currentRoute.substring(3); 
            console.log("Remaining part:", remainingPart);
            setUsername(remainingPart)
        } else {
            console.log("Current route doesn't start with '/u/'");
        }
                
    },[])

    React.useEffect(()=>{
        async function fetch(){
            let resp = await GetStudentData(username)
            setData(resp)
        }
        fetch()
    },[username])

    return (
        <>
            <h1>Hey! {username}</h1>
            {data && (
                <div>
                    <p>ID: {data.id}</p>
                    <p>Login: {data.login}</p>
                    <p>Email: {data.email}</p>
                    <p>Klass: {data.klass}</p>
                    <p>School: {data.school}</p>
                    <p>Stats:</p>
                    <ul>
                        <li>Solved: {data.stats.solved}</li>
                        <li>LeetCode: {data.stats.leetcode}</li>
                        <li>Badges: {data.stats.badges}</li>
                        <li>Rating: {data.stats.rating}</li>
                        <li>Rank: {data.stats.rank}</li>
                    </ul>
                </div>
            )}
        </>
    )
}

export default Profile;
