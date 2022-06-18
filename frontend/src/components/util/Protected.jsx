import React, { useContext } from 'react';
import { Navigate } from 'react-router-dom';
import { AuthContext } from '../../context/auth';

function Protected({ name, children }) {
  const {context} = useContext(AuthContext);
  const { allowed } = context;

  return (
    <>
      {allowed === true ? children : <Navigate to='/' />}
    </>
  )
}

export default Protected;