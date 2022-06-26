import React, { useCallback, useEffect, useState } from 'react';
import { verifyToken } from '../api/Auth';

export const AuthContext = React.createContext();

export async function checkToken() {
  let token = localStorage.getItem('token');

  if (token == null || token === '') {
    return { error: true };
  }

  let resp = await verifyToken(token);
  if (!resp.error) {
    localStorage.setItem('token', resp.token);
    sessionStorage.setItem('authenticated', JSON.stringify(true));
  }

  return resp; 
}

export function AuthProvider({ children }) {
  const [context, setState] = useState({});

  const getContext = useCallback(async () => {
    if (context.token !== '' && localStorage.getItem('token') === null) {
      localStorage.setItem('token', context.token);
    }

    let resp = await checkToken();
    if (!resp.error) {
      setToken(resp.token);
      return;
    }

    setError({ error: true, msg: 'Failed to authenticate user' });
  }, []);

  const setError = (error) => {
    setState({ ...context, error })
  }

  useEffect(() => {
    getContext()
  }, [])
  
  const clearAuth = () => {
    setState({
      ...context,
      allowed: false,
      token: '',
      payload: null,
    });
  };

  const setToken = (token) => {
    localStorage.setItem('token', token);

    const parts = token.split('.');
    const data = JSON.parse(atob(parts[1]));

    setState({
      ...context,
      allowed: true,
      token,
      payload: data,
    });
  };

  const check = async (p) => {
    let resp = await p;
    if (resp.unauthorized === true) {
      clearAuth()
      return { error: true, message: "Unauthorized" };
    }

    return resp;
  };

  return (
    <AuthContext.Provider value={{ context, getContext, clearAuth, setToken, check }}>
      {children}
    </AuthContext.Provider>
  )
}

