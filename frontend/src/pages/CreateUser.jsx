import React, { useState } from 'react';
import { createUser } from '../api';
import { useNavigate } from 'react-router-dom';

export default function CreateUser() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    await createUser({ Name: name, Email: email });
    navigate('/');
  };

  return (
    <form onSubmit={handleSubmit}>
      <div className="mb-3">
        <label>Name</label>
        <input type="text" className="form-control" value={name} onChange={(e) => setName(e.target.value)} required />
      </div>
      <div className="mb-3">
        <label>Email</label>
        <input type="email" className="form-control" value={email} onChange={(e) => setEmail(e.target.value)} required />
      </div>
      <button className="btn btn-success">Create User</button>
    </form>
  );
}
