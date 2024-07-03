// src/TokenForm.js
import React, { useState } from 'react';
import axios from 'axios';
import './TokenForm.css';

const TokenForm = ({ setToken }) => {
  const [name, setUsername] = useState('');
  //const [password, setPassword] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try  {
      const response = await axios.get(
        'http://localhost:8080/token?name=' + name, 
        {
          headers: {
            'Content-Type': 'application/json',
          },
        }
      );
      setToken(response.data);
    } catch (error) {
      console.error('Error fetching the token:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        placeholder="Username"
        value={name}
        onChange={(e) => setUsername(e.target.value)}
        required
      />
      {/* <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        required
      /> */}
      <button type="submit">Get Token</button>
    </form>
  );
};

export default TokenForm;
