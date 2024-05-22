export async function GetTests(username) {
    try {
        const response = await fetch(`http://localhost:8080/tests/${username}`);
        console.log(response);
        const data = await response.json();
        if (data && data.length > 0) {
            data[0].input = data[0].input.replace(/\n/g, '\\n');
            data[0].output = data[0].output.replace(/\n/g, '\\n');
            return data[0];
        } else {
            throw new Error('No test data found');
        }
    } catch (error) {
        console.error('Error fetching classwork:', error);
        throw error;
    }
}
