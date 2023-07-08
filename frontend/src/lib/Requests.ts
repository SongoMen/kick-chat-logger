const makeAPIRequest = async (path: string) => {
    const response = await fetch("http://localhost:8080/api/v1" + path);
    const data = await response.json();
    return data;
};

export {
    makeAPIRequest
}