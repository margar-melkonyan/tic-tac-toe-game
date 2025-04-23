import axios from "axios";

const http = axios.create({
  baseURL: 'https://your-api-url.com/api',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Добавь interceptors если нужно (например, токен)
http.interceptors.request.use(config => {
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export default http;
