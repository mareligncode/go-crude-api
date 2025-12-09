import React from 'react';
import { Routes, Route, Link } from 'react-router-dom';
import UserList from './pages/UserList';
import CreateUser from './pages/CreateUser';
import EditUser from './pages/EditUser';

export default function App() {
  return (
    <div className="container mt-5">
      <nav className="mb-4">
        <Link to="/" className="btn btn-primary me-2">Users</Link>
        <Link to="/create" className="btn btn-success">Create User</Link>
      </nav>
      <Routes>
        <Route path="/" element={<UserList />} />
        <Route path="/create" element={<CreateUser />} />
        <Route path="/edit/:id" element={<EditUser />} />
      </Routes>
    </div>
  );
}
