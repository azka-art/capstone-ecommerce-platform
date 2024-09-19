import { fetchFromApi } from './apiservice';

export const getCartItems = () => {
  return fetchFromApi('cart');
};

export const addToCart = (product) => {
  return fetchFromApi('cart', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(product),
  });
};

export const removeFromCart = (id) => {
  return fetchFromApi(`cart/${id}`, {
    method: 'DELETE',
  });
};
