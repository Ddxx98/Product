// src/App.js
import React, { useState } from 'react';
import TokenForm from './components/token/TokenForm';
import ProductForm from './components/productForm/ProductForm';
import ProductList from './components/productList/ProductList';
import './App.css';

const App = () => {
  const [reload, setReload] = useState(false);
  const [token, setToken] = useState('');

  const handleProductAdded = () => {
    setReload((prev) => !prev); // Toggle the reload state to trigger re-fetch
  };

  return (
    <div>
      <h1>Product Management</h1>
      {!token ? (
        <TokenForm setToken={setToken} />
      ) : (
        <>
          <ProductForm token={token}  onProductAdded={handleProductAdded} />
          <ProductList token={token} reload={reload} />
        </>
      )}
    </div>
  );
};

export default App;
