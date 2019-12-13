import axios from 'axios';

var axiosInstance = axios.create({
  baseURL: 'http://localhost:8080/',
});

export default axiosInstance