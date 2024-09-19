import React from 'react';

const CartItem = ({ item, onRemove }) => {
  return (
    <div className="cart-item">
      <h3>{item.product.name}</h3>
      <p>Quantity: {item.quantity}</p>
      <p>${item.product.price * item.quantity}</p>
      <button onClick={() => onRemove(item)}>Remove</button>
    </div>
  );
};

export default CartItem;

