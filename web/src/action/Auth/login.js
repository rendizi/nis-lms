export async function LoginStudentAPICall(formData){
    try {
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        });

        if (response.status === 200) {
            const data = await response.json()
            return { code: 200, token: data.token,message:"success" };
        } else {
            const errorData = await response.json();
            return { code: response.status, message: errorData.message || "An error occurred" };
        }
    } catch (error) {
        return { code: 400, message: error.message || "Network error" };
    }
}


export async function LoginTeacherAPICall(formData){
    try {
        const response = await fetch('http://localhost:8080/login?teacher=1', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(formData)
        });

        if (response.status === 200) {
            const data = await response.json()
            return { code: 200, token: data.token,message:"success" };
        } else {
            const errorData = await response.json();
            return { code: response.status, message: errorData.message || "An error occurred" };
        }
    } catch (error) {
        return { code: 400, message: error.message || "Network error" };
    }
}


