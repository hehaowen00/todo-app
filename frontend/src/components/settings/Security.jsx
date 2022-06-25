import React, { useContext, useState } from 'react';
import { Alert, Button, Form, Row, Col, Tab } from 'react-bootstrap';

import { AuthContext } from '../../context/Auth';
import { updatePassword } from '../../api/User';

function Security() {
  const { context, check } = useContext(AuthContext);

  const [state, setState] = useState({
    error: '',
    success: '',
    password: '',
    new_password: '',
  });

  const updateField = (e) => {
    const { name, value } = e.target;
    setState({ ...state, [name]: value });
  };

  const updatePass = async (e) => {
    e.preventDefault();
    const { token } = context;
    const { password, new_password } = state;
    const resp = await check(updatePassword(token, password, new_password));

    if (resp.error) {
      setState({
        password: '',
        new_password: '',
        error: resp.message,
        success: '',
      });
      return;
    }

    setState({
      password: '',
      new_password: '',
      error: '',
      success: 'Password updated',
    });
  };

  return (
    <Tab.Pane eventKey="security">
      <h5>Change Password</h5>
      <hr />
      {state.success !== '' && 
        <Alert
          variant='success'
        >
          {state.success}
        </Alert>
      }
      <Form onSubmit={updatePass}>
        <Row>
          <Form.Group as={Col} md='12'>
            <Form.Label>Current Password</Form.Label>
            <Form.Control
              className='mb-3'
              type='password'
              placeholder='Current Password'
              name='password'
              onChange={updateField}
              value={state.password}
              required
            />
            <Form.Label>New Password</Form.Label>
            <Form.Control
              className='mb-3'
              type='password'
              placeholder='New Password'
              name='new_password'
              onChange={updateField}
              value={state.new_password}
              required
            />
          </Form.Group>
        </Row>
        {state.error !== '' && 
          <Alert
            variant='danger'
          >
            {state.error}
          </Alert>
        }
        <Row className='justify-content-end'>
          <Col md='4' sm='6' xs='12'>
            <Button
              className='w-100'
              type='submit'
              variant='primary'
            >
              Change Password
            </Button>
          </Col>
        </Row>
      </Form>
    </Tab.Pane>
  )
}

export default Security;
