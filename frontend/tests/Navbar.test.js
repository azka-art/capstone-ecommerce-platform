import React from 'react';
import { render, screen } from '@testing-library/react';
import Navbar from '../components/Navbar';

test('renders Navbar component', () => {
  render(<Navbar />);
  expect(screen.getByText(/Home/i)).toBeInTheDocument();
  expect(screen.getByText(/Products/i)).toBeInTheDocument();
  expect(screen.getByText(/Cart/i)).toBeInTheDocument();
});
