export async function IsItSolved(username, id) {
    try {
        const response = await fetch(`http://localhost:8080/u/${username}/${id}`);
        if (response.status !== 200) {
            return false;
        }
        const respData = await response.json();
        return respData.did;
    } catch (error) {
        console.error("Error:", error);
        return false;
    }
}
