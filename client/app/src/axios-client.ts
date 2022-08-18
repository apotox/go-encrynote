import axios from "axios";
import getConfig from "./config";

const Client = axios.create({
  baseURL: getConfig().apiUrl,
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
  },
});


export default Client;
