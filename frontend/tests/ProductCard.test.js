import React from 'react';
import { render, screen } from '@testing-library/react';
import ProductCard from '../components/ProductCard';

const product = {
  name: 'Test Product',
  description: 'Test Description',
  price: 99.99,
};

test('renders ProductCard component', () => {
  render(<ProductCard product={product} />);
  expect(screen.getByText(/Test Product/i)).toBeInTheDocument();
  expect(screen.getByText(/Test Description/i)).toBeInTheDocument();
  expect(screen.getByText(/\$99\.99/i)).toBeInTheDocument();
});
