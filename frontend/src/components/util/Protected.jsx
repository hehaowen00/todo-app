import React, { useContext, useEffect } from 'react';
import { Container, Spinner } from 'react-bootstrap';
import { Navigate } from 'react-router-dom';

import Topbar from './Topbar';

import { AuthContext } from '../../context/Auth';

function Protected({ children }) {
  const { context, getContext } = useContext(AuthContext);

  useEffect(() => {
    if (Object.keys(context).length === 0) {
      getContext();
    }
  }, []);

  if (Object.keys(context).length === 0) {
    return (
      <>
        <Topbar />
        <Container>
          <Spinner animation="border" role="status">
            <span className="visually-hidden">Loading...</span>
          </Spinner>
        </Container>
      </>
    )
  }

  const { allowed } = context;

  return (
    <>
      { allowed === true ? children : <Navigate to='/' />}
    </>
  )
}

export default Protected;