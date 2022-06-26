import React, { useCallback, useContext, useEffect, useState } from 'react';
import { Alert, Button, Col, Container, Form, Row } from 'react-bootstrap'
import { useNavigate } from 'react-router-dom';

import Topbar from './util/Topbar';

import { AuthContext, checkToken } from '../context/Auth';
import { loginUser } from '../api/Auth'
import { setTitle } from '../util/Util';

function Login() {
  const { context, clearAuth, setToken } = useContext(AuthContext);

  let [state, setState] = useState({ error: '', username: '', password: '', });

  let nav = useNavigate();

  const updateField = (e) => {
    e.preventDefault();
    setState({ ...state, [e.target.name]: e.target.value });
  };

  const updateString = (e) => {
    setState({ ...state, [e.target.name]: e.target.value.trim() });
  };

  const login = async (e) => {
    e.preventDefault();
    setState({ ...state, error: ''});

    let resp = await loginUser(state.username, state.password);

    if (resp.error) {
      setState({ ...state, error: resp.message });
      clearAuth();
      return
    }

    setToken(resp.token);
    nav('/home');
  };

  const verify = useCallback(async () => {
    let resp = await checkToken();
    if (!resp.error) {
      setToken(resp.token);
      nav('/home');
    }
  }, []);

  useEffect(() => {
    setTitle('Login');
    verify();
  }, [context, nav, verify]);

  return (
    <>
      <Topbar />
      <Container>
        <br />
        <Row className='justify-content-center'>
          <Col className='text-center' md={6}>
            <h3>Login</h3>
          </Col>
        </Row>
        <Row className='justify-content-center'>
          <Col md={5}>
            <Form onSubmit={login}>
              <Form.Group className='mb-3'>
                <Form.Label>Username</Form.Label>
                <Form.Control
                 name='username'
                 type='username'
                 placeholder='Username'
                 onChange={updateString}
                 value={state.username}
                 required
                />
              </Form.Group>
              <Form.Group className='mb-3'>
                <Form.Label>Password</Form.Label>
                <Form.Control
                 name='password'
                 type='password'
                 placeholder='Password'
                 onChange={updateField}
                 value={state.password}
                 required
              />
              </Form.Group>
              {state.error !== '' && 
                <Alert variant='danger'>
                    {state.error}
                </Alert>
              }

              <p className='text-end'>
              <a href='/register'>Don't have an account?</a>
              </p>
              <Button
                className='w-100'
                type='submit'
                variant='primary'
              >
                Login
              </Button>
            </Form>
          </Col>
        </Row>
      </Container>
    </>
  )
}

export default Login;
