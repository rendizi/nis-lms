export async function FetchStudentsClassWork(username) {
    try {
        const response = await fetch(`http://localhost:8080/u/${username}/classwork`);
        console.log(response)
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching classwork:', error);
        throw error;
    }
}

export async function FetchTeachersClassWork(username) {
    try {
        const response = await fetch(`http://localhost:8080/t/${username}/classwork`);
        console.log(response)
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching classwork:', error);
        throw error;
    }
}
