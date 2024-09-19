import React from 'react';
import { render, screen } from '@testing-library/react';
import LoginPage from '../pages/LoginPage';

test('renders LoginPage component', () => {
  render(<LoginPage />);
  expect(screen.getByLabelText(/Username/i)).toBeInTheDocument();
  expect(screen.getByLabelText(/Password/i)).toBeInTheDocument();
  expect(screen.getByRole('button', { name: /Login/i })).toBeInTheDocument();
});

