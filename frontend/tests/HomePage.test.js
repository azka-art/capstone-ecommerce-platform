import React from 'react';
import { render, screen } from '@testing-library/react';
import HomePage from '../pages/HomePage';

test('renders HomePage component', () => {
  render(<HomePage />);
  expect(screen.getByText(/Welcome to Our Shop/i)).toBeInTheDocument();
  expect(screen.getByRole('link', { name: /Products/i })).toBeInTheDocument();
});
