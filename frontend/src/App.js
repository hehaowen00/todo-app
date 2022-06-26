import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import Home from './components/Home';
import Register from './components/Register';
import Login from './components/Login';
import Settings from './components/Settings';

import List from './components/todos/list/List';
import NewList from './components/todos/NewList';
import Todo from './components/todos/Todo';

import Logout from './components/util/Logout';
import NotFound from './components/util/NotFound';
import Protected from './components/util/Protected';

import { AuthProvider } from './context/Auth';
import { ListProvider } from './context/List';

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
          <Route exact path='/register' element={<Register />} />
          <Route exact path='/logout' element={<Logout />} />
          <Route exact path='/home' element={<Protected><Home /></Protected>} />
          <Route exact path='/list/new' element={<Protected><NewList/></Protected>} />
          <Route path='/list/:id' element={<Protected><List /></Protected>} />
          <Route path='/item/:id' element={<Protected><Todo /></Protected>} />
          <Route path='/settings' element={<Protected><Settings /></Protected>} />
        </Routes>
      </BrowserRouter>
      </ListProvider>
    </AuthProvider>
  );
}

export default App;
