export function isTokenValid(token) {
    try {
        const parts = token.split('.');
        if (parts.length !== 3) {
            throw new Error('Invalid token format');
        }

        const payload = JSON.parse(atob(parts[1]));
        const expirationTime = payload.exp * 1000; 
        const currentTime = Date.now();

        return currentTime <= expirationTime;
    } catch (error) {
        console.error('Error decoding or validating token:', error);
        return false; 
    }
}