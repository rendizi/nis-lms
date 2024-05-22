export async function GetTaskData(username) {
    try {
        const response = await fetch(`http://localhost:8080/task/${username}`);
        console.log(response)
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching classwork:', error);
        throw error;
    }
}
