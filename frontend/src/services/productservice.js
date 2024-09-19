import { fetchFromApi } from './apiservice';

export const getProducts = () => {
  return fetchFromApi('products');
};

export const getProductById = (id) => {
  return fetchFromApi(`products/${id}`);
};
