import React, { useContext, useState } from 'react';
import { Alert, Button, Form, Row, Col, Tab } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';

import { AuthContext } from '../../context/Auth';
import { deleteAccount } from '../../api/User';

function DeleteAccount() {
  const { context, clearAuth, check } = useContext(AuthContext);
  const nav = useNavigate();

  const [state, setState] = useState({
    password: '',
    confirmation: '',
    error: '',
  });

  const confirmText = 'Delete My Account';

  const updateState = (e) => {
    const { name, value } = e.target;
    setState({ ...state, [name]: value });
  };

  const deleteUser  = async (e) => {
    e.preventDefault();

    const { token } = context;
    const { password } = state;

    const resp = await check(deleteAccount(token, password));
    if (resp.error) {
      setState({ ...state, error: resp.message });
      return;
    }

    clearAuth();
    nav('/');
  };

  return (
    <Tab.Pane eventKey="delete">
      <h5>Delete Account</h5>
      <hr />
      <Form onSubmit={deleteUser}>
        <Row>
          <Form.Group as={Col} md='12'>
            <Form.Label>Current Password</Form.Label>
            <Form.Control
              className='mb-3'
              type='password'
              name='password'
              placeholder='Current Password'
              onChange={updateState}
              value={state.password}
              required
            />
          </Form.Group>
        </Row>
        <Row>
          <Form.Group as={Col} md='12'>
            <Form.Label>Please type "{confirmText}" to confirm</Form.Label>
            <Form.Control
              className='mb-3'
              type='input'
              name='confirmation'
              onChange={updateState}
              value={state.confirmation}
              required
            />
          </Form.Group>
        </Row>
        { state.error &&
          <Alert variant='danger'>
            {state.error}
          </Alert>
        }
        <Row className='justify-content-end'>
          <Col md='3' xs='12'>
            <Button
              className='w-100'
              type='submit'
              variant='danger'
              disabled={state.confirmation !== confirmText}
            >
              Delete Account
            </Button>
          </Col>
        </Row>
      </Form>
    </Tab.Pane>
  )
}

export default DeleteAccount;
