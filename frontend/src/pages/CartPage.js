import React, { useState, useEffect } from 'react';
import { getCartItems, removeFromCart } from '../services/cartService';
import CartItem from '../components/cartitem';

const CartPage = () => {
  const [cartItems, setCartItems] = useState([]);

  useEffect(() => {
    const fetchCartItems = async () => {
      const data = await getCartItems();
      setCartItems(data);
    };

    fetchCartItems();
  }, []);

  const handleRemove = async (item) => {
    await removeFromCart(item.id);
    setCartItems(cartItems.filter(ci => ci.id !== item.id));
  };

  return (
    <div>
      <h1>Your Cart</h1>
      {cartItems.length === 0 ? (
        <p>Your cart is empty.</p>
      ) : (
        <div>
          {cartItems.map(item => (
            <CartItem key={item.id} item={item} onRemove={handleRemove} />
          ))}
        </div>
      )}
    </div>
  );
};

export default CartPage;
