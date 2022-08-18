


interface Config {
    apiUrl: string;
    REACT_APP_APP_SHORTSHA: string;
    REACT_APP_APP_BRANCH: string;
    REACT_APP_URL: string;
}

function getConfig(): Config{

    return {
        apiUrl: process.env.REACT_APP_API_URL ?? 'http://localhost:3000/api/',
        REACT_APP_APP_SHORTSHA: process.env.REACT_APP_APP_SHORTSHA ?? 'v1',
        REACT_APP_APP_BRANCH: process.env.REACT_APP_SHA_BRANCH ?? 'master',
        REACT_APP_URL: process.env.REACT_APP_URL ?? 'http://localhost:3000',
    }

}

export default getConfig;