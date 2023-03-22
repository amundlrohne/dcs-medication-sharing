export const resolveURL = (service) => {
    const prod = process.env.REACT_APP_PRODUCTION;
    if (prod === "true") return "";

    switch (service) {
        case "consent":
            return "http://localhost:8180";
        case "healthcare-provider":
            return "http://localhost:8280";
        case "medication-record":
            return "http://localhost:8380";
        case "standardization":
            return "http://localhost:8480";
        default:
            return "http://localhost";
    }
};
