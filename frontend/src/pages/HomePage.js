import React, { useEffect, useState } from 'react';
import { getProducts } from '../services/productService';
import ProductCard from '../components/productcard';

const HomePage = () => {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    const fetchProducts = async () => {
      const data = await getProducts();
      setProducts(data);
    };

    fetchProducts();
  }, []);

  const handleAddToCart = (product) => {
    // Logic for adding product to cart
  };

  return (
    <div>
      <h1>Welcome to our E-Commerce Platform</h1>
      <div className="product-list">
        {products.map(product => (
          <ProductCard
            key={product.id}
            product={product}
            onAddToCart={handleAddToCart}
          />
        ))}
      </div>
    </div>
  );
};

export default HomePage;

