import { fetchFromApi } from './apiservice';

export const register = (username, password) => {
  return fetchFromApi('register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username, password }),
  });
};

export const login = (username, password) => {
  return fetchFromApi('login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username, password }),
  });
};

export const getUserProfile = () => {
  const token = localStorage.getItem('token');
  return fetchFromApi('profile', {
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });
};

export const isAuthenticated = () => {
    const token = localStorage.getItem('token');
    // Example token validation logic
    return token != null && token !== '';
  };
  

