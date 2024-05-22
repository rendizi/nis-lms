export async function RegisterStudentAPICall(formData){
    try {
        const response = await fetch('http://localhost:8080/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        });

        if (response.status === 200) {
            return { code: 200, message: "success" };
        } else {
            const errorData = await response.json();
            return { code: response.status, message: errorData.message || "An error occurred" };
        }
    } catch (error) {
        return { code: 400, message: error.message || "Network error" };
    }
}


