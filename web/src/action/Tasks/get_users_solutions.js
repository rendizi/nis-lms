export async function GetStudentSolutions(username, page) {
    try {
        const url = page
            ? `http://localhost:8080/u/${username}/solutions?page=${page}`
            : `http://localhost:8080/u/${username}/solutions`;

        const response = await fetch(url);
        console.log(response);

        if (!response.ok) {
            throw new Error('Network response was not ok');
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching solutions:', error);
        throw error;
    }
}
