import React, { useCallback, useContext, useEffect } from 'react';
import { Navigate } from 'react-router-dom';
import { AuthContext, checkToken } from '../../context/auth';

function Protected({ children }) {
  const { context, setToken } = useContext(AuthContext);

  const verify = useCallback(async () => {
    let resp = await checkToken();
    if (!resp.error) {
      setToken(resp.token);
    }
  }, []);

  useEffect(() => {
    if (Object.keys(context).length === 0) {
      verify();
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