// src/ProductList.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './ProductList.css';

const ProductList = ({ token, reload }) => {
  const [products, setProducts] = useState([]);

  const DeleteProduct = async (id) => {
    try {
      const response = await axios.delete(`http://localhost:8080/products/${id}`, {
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      console.log(response.data);
    } catch (error) {
      console.error('Error deleting product:', error);
    }
  };

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const response = await axios.get('http://localhost:8080/products', {
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${token}`,
          },
        });
        setProducts(response.data);
      } catch (error) {
        console.error('Error fetching products:', error);
      }
    };

    fetchProducts();
  }, [token, reload, DeleteProduct]);

  return (
    <div>
      <h2>Product List</h2>
      <ul>
        {products.length != 0 ? (products.map((product) => (
          <li key={product.id} >{product.name} - ${product.price} <button onClick={() => DeleteProduct(product.id)}>Delete</button> </li>
        ))):( <li>No products found</li>)
        }
      </ul>
    </div>
  );
};

export default ProductList;
