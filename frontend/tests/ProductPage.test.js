import React from 'react';
import { render, screen } from '@testing-library/react';
import ProductPage from '../pages/ProductPage';

const product = {
  name: 'Test Product',
  description: 'Test Description',
  price: 99.99,
};

test('renders ProductPage component', () => {
  render(<ProductPage product={product} />);
  expect(screen.getByText(/Test Product/i)).toBeInTheDocument();
  expect(screen.getByText(/Test Description/i)).toBeInTheDocument();
  expect(screen.getByText(/\$99\.99/i)).toBeInTheDocument();
});
