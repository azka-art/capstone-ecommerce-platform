import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import { getProductById } from '../services/productService';
import ProductCard from '../components/productcard';

const ProductPage = () => {
  const { id } = useParams();
  const [product, setProduct] = useState(null);

  useEffect(() => {
    const fetchProduct = async () => {
      const data = await getProductById(id);
      setProduct(data);
    };

    fetchProduct();
  }, [id]);

  if (!product) return <p>Loading...</p>;

  const handleAddToCart = () => {
    // Logic for adding product to cart
  };

  return (
    <div>
      <ProductCard product={product} onAddToCart={handleAddToCart} />
    </div>
  );
};

export default ProductPage;

