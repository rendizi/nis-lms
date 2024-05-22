export async function GetStudentData(username) {
    try {
        const response = await fetch(`http://localhost:8080/u/${username}`);
        console.log(response)
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching classwork:', error);
        throw error;
    }
}
