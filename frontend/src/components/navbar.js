import React from 'react';
import { Link } from 'react-router-dom';
import { Button } from './Button';

const Navbar = () => {
  return (
    <nav>
      <h1>My E-Commerce Platform</h1>
      <ul>
        <li><Link to="/">Home</Link></li>
        <li><Link to="/products">Products</Link></li>
        <li><Link to="/cart">Cart</Link></li>
        <li><Link to="/profile">Profile</Link></li>
        <li><Link to="/login"><Button>Login</Button></Link></li>
      </ul>
    </nav>
  );
};

export default Navbar;
