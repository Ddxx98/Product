// src/ProductForm.js
import React, { useState } from 'react';
import axios from 'axios';
import './ProductForm.css';

const ProductForm = ({ token, onProductAdded }) => {
  const [productName, setProductName] = useState('');
  const [productPrice, setProductPrice] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post(
        'http://localhost:8080/products',
        { Name: productName, Price: Number(productPrice) },
        {
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
          },
        }
      );
      setProductName('');
      setProductPrice('');
      onProductAdded();
    } catch (error) {
      console.error('Error adding the product:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        placeholder="Product Name"
        value={productName}
        onChange={(e) => setProductName(e.target.value)}
        required
      />
      <input
        type="number"
        placeholder="Product Price"
        value={productPrice}
        onChange={(e) => setProductPrice(e.target.value)}
        required
      />
      <button type="submit">Add Product</button>
    </form>
  );
};

export default ProductForm;
