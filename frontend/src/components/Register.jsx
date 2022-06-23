import React, { useCallback, useContext, useEffect, useState } from 'react';
import { Alert, Button, Col, Container, Form, Row } from 'react-bootstrap'
import { Link, useNavigate } from 'react-router-dom';

import Topbar from './util/Topbar';

import { AuthContext, checkToken } from '../context/Auth';
import { registerUser } from '../api/Auth';
import { setTitle } from '../util/Util';

function Register() {
  const { setToken } = useContext(AuthContext);
  const nav = useNavigate();

  let [state, setState] = useState({
    error: '',
    success: '',
    username: '',
    password: '',
  });

  const updateUsername = (e) => {
    setState({ ...state, username: e.target.value.trim() });
  };

  const updatePassword = (e) => {
    setState({ ...state, password: e.target.value });
  };

  const register = async (e) => {
    e.preventDefault();
    setState({ ...state, error: '', success: '' });

    let resp = await registerUser(state.username, state.password);
    let message = resp.message;

    if (resp.error) {
      setState({ ...state, error: message });
      return
    }

    setState({ error: '', success: message, username: '', password: '', });
  };

  const verify = useCallback(async () => {
    let resp = await checkToken();
    if (!resp.error) {
      setToken(resp.token);
      nav('/home');
    }
  }, [setToken, nav]);

  useEffect(() => {
    setTitle('Register');
    verify();
  }, [verify]);

  return (
    <>
      <Topbar />
      <Container>
        <br />
        <Row className='justify-content-center'>
          <Col className='text-center' md={6}>
            <h3>Register</h3>
          </Col>
        </Row>
        <Row className='justify-content-center'>
          <Col md={5}>
            <Form onSubmit={register}>
              <Form.Group className='mb-3'>
                <Form.Label>Username</Form.Label>
                <Form.Control
                 type='username'
                 placeholder='Username'
                 onChange={updateUsername}
                 value={state.username}
                 required
                />
              </Form.Group>
              <Form.Group className='mb-3'>
                <Form.Label>Password</Form.Label>
                <Form.Control
                 type='password'
                 placeholder='Password'
                 onChange={updatePassword}
                 value={state.password}
                 required
                />
              </Form.Group>
              {state.error && 
                <Alert variant='danger'>
                  {state.error}
                </Alert>
              }
              {state.success &&
                <Alert variant='success'>
                  User Created. <span>Click <Link to='/'>Here</Link> to login.</span>
                </Alert>
              }
              <Button
               className='full-width'
               type='submit'
               variant='primary'
              >
                Register
              </Button>
            </Form>
          </Col>
        </Row>
      </Container>
    </>
  );
}

export default Register;
