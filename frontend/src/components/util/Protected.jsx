import React, { useContext, useEffect } from 'react';
import { Navigate } from 'react-router-dom';
import { AuthContext } from '../../context/Auth';

function Protected({ children }) {
  const { context, getContext } = useContext(AuthContext);

  useEffect(() => {
    if (Object.keys(context).length === 0) {
      getContext()
    }
  }, []);

  if (Object.keys(context).length === 0) {
    return (<></>);
  }

  const { allowed } = context;

  return (
    <>
      {allowed === true ? children : <Navigate to='/' />}
    </>
  )
}

export default Protected;