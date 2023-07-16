const makeAPIRequest = async (path: string) => {
    const response = await fetch(import.meta.env.VITE_API_BASE_URL + path);
    const data = await response.json();
    return data;
};

export {
    makeAPIRequest
}