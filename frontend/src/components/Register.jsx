import React, { useCallback, useContext, useEffect, useState } from 'react';
import { Alert, Button, Col, Container, Form, Row } from 'react-bootstrap'
import { Link, useNavigate } from 'react-router-dom';

import Topbar from './util/Topbar';

import { AuthContext, checkToken } from '../context/Auth';
import { registerUser } from '../api/Auth';

function Register() {
  const { setToken } = useContext(AuthContext);
  const nav = useNavigate();

  let [state, setState] = useState({});
  let [error, setError] = useState('');
  let [success, setSuccess] = useState('');

  const updateUsername = (e) => {
    setState({ ...state, username: e.target.value.trim() });
  };

  const updatePassword = (e) => {
    setState({ ...state, password: e.target.value });
  };

  const register = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');

    let resp = await registerUser(state.username, state.password);

    if (resp.error) {
      setError(resp.message);
      return
    }

    setSuccess('User Created.');
    setState({});
  };

  const verify = useCallback(async () => {
    let resp = await checkToken();
    if (!resp.error) {
      setToken(resp.token);
      nav('/home');
    }
  }, [setToken, nav]);

  useEffect(() => {
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
                 name='username'
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
                 name='password'
                 type='password'
                 placeholder='Password'
                 onChange={updatePassword}
                 value={state.password}
                 required
                />
              </Form.Group>
              {error && 
                <Alert variant='danger'>
                  {error}
                </Alert>
              }
              {success &&
                <Alert variant='success'>
                  {success} <span>Click <Link to='/'>Here</Link> to login.</span>
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