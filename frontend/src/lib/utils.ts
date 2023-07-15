const padNumber = (number: number, digits: number) => {
    return String(number).padStart(digits, '0');
};

export {
    padNumber
}