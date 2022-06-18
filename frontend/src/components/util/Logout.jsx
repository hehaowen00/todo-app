import React, { useContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../../context/Auth';

function Logout() {
  const { clearAuth } = useContext(AuthContext);

  const nav = useNavigate();

  useEffect(() => {
    localStorage.removeItem('token');
    clearAuth();
    nav('/');
  }, [clearAuth, nav]);

  return (
    <></>
  )
}

export default Logout;