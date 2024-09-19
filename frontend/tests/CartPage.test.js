import React from 'react';
import { render, screen } from '@testing-library/react';
import CartPage from '../pages/CartPage';

const cart = [
  { name: 'Test Product', quantity: 1, price: 99.99 },
];

test('renders CartPage component', () => {
  render(<CartPage cart={cart} />);
  expect(screen.getByText(/Test Product/i)).toBeInTheDocument();
  expect(screen.getByText(/Quantity: 1/i)).toBeInTheDocument();
  expect(screen.getByText(/\$99\.99/i)).toBeInTheDocument();
});
