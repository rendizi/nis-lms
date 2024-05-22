export async function SearchTask(props) {
    let url = 'http://localhost:8080/search/task';
    
    let queryParams = [];
    queryParams.push(`pagesize=8`);
    if (props.page !== null && props.page !== undefined && props.page >= 0){
        queryParams.push(`page=${props.page}`);
    }
    if (props.difficulty !== null && props.difficulty !== undefined && props.difficulty != "") {
        queryParams.push(`difficulty=${encodeURIComponent(props.difficulty)}`);
    }
    if (props.title !== null && props.title !== undefined && props.title != "") {
        queryParams.push(`title=${encodeURIComponent(props.title)}`);
    }
    if (queryParams.length > 0) {
        url += '?' + queryParams.join('&');
    }
    try {
        const response = await fetch(url);
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error('Error fetching data:', error);
        throw error;
    }
}
