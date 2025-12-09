import React, { useEffect, useState } from 'react';
import { getUser, updateUser } from '../api';
import { useNavigate, useParams } from 'react-router-dom';

export default function EditUser() {
  const { id } = useParams();
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    fetchUser();
  }, []);

  const fetchUser = async () => {
    const res = await getUser(id);
    setName(res.data.Name);
    setEmail(res.data.Email);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    await updateUser(id, { Name: name, Email: email });
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
      <button className="btn btn-primary">Update User</button>
    </form>
  );
}
