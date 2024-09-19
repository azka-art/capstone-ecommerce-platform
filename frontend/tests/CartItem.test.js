import React from 'react';
import { render, screen } from '@testing-library/react';
import CartItem from '../components/CartItem';

const cartItem = {
  name: 'Test Product',
  quantity: 2,
  price: 99.99,
};

test('renders CartItem component', () => {
  render(<CartItem item={cartItem} />);
  expect(screen.getByText(/Test Product/i)).toBeInTheDocument();
  expect(screen.getByText(/Quantity: 2/i)).toBeInTheDocument();
  expect(screen.getByText(/\$99\.99/i)).toBeInTheDocument();
});
