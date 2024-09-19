import React from 'react';
import { render, screen } from '@testing-library/react';
import ProfilePage from '../pages/ProfilePage';

const user = {
  username: 'testuser',
};

test('renders ProfilePage component', () => {
  render(<ProfilePage user={user} />);
  expect(screen.getByText(/Username: testuser/i)).toBeInTheDocument();
});
