import axios from 'axios';

let path = window.location.pathname;

if (path.substr(path.length - 7, 7) === "/katib/") {
  path = path.substr(0, path.length - 7)
}

var axiosInstance = axios.create({
  baseURL: window.location.origin + path,
});

export default axiosInstance