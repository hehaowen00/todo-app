import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import Home from './components/Home';
import Register from './components/Register';
import Login from './components/Login';

import List from './components/todos/List';
import NewList from './components/todos/NewList';
import Todo from './components/todos/Todo';

import Logout from './components/util/Logout';
import NotFound from './components/util/NotFound';
import Protected from './components/util/Protected';

import { AuthProvider } from './context/auth';
import { ListProvider } from './context/list';

import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';

function App() {
  return (
    <AuthProvider>
      <ListProvider>
      <BrowserRouter>
        <Routes>
          <Route path='*' element={<NotFound />} />
          <Route exact path='/' element={<Login />} />
          <Route path='/register' element={<Register />} />
          <Route path='/logout' element={<Logout />} />
          <Route path='/home' element={<Protected><Home /></Protected>} />
          <Route path='/newList' element={<Protected><NewList/></Protected>} />
          <Route path='/list/:id' element={<Protected><List /></Protected>} />
          <Route path='/item/:id' element={<Protected><Todo /></Protected>} />
        </Routes>
      </BrowserRouter>
      </ListProvider>
    </AuthProvider>
  );
}

export default App;