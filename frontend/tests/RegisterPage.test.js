import React from 'react';
import { render, screen } from '@testing-library/react';
import RegisterPage from '../pages/RegisterPage';

test('renders RegisterPage component', () => {
  render(<RegisterPage />);
  expect(screen.getByLabelText(/Username/i)).toBeInTheDocument();
  expect(screen.getByLabelText(/Password/i)).toBeInTheDocument();
  expect(screen.getByLabelText(/Confirm Password/i)).toBeInTheDocument();
  expect(screen.getByRole('button', { name: /Register/i })).toBeInTheDocument();
});
