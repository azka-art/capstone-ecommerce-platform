const API_URL = 'http://localhost:5000';

export const fetchFromApi = async (endpoint, options = {}) => {
  const response = await fetch(`${API_URL}/${endpoint}`, options);
  if (!response.ok) {
    throw new Error('Network response was not ok');
  }
  return response.json();
};

