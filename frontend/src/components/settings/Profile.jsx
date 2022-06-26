import React, { useContext, useState } from 'react';
import { Alert, Button, Form, Row, Col, Tab } from 'react-bootstrap';

import { AuthContext } from '../../context/Auth';
import { DownloadData, updateDetails } from '../../api/User';

function Profile() {
  const { context, check, setToken } = useContext(AuthContext);
  const { username } = context.payload;

  const [state, setState] = useState({
    username,
    password: '',
    error: '',
    success: '',
  });

  const updateState = (e) => {
    const { name, value } = e.target;
    setState({ ...state, [name]: value });
  };

  const updateProfile = async (e) => {
    e.preventDefault();

    const { token } = context;
    const { username, password } = state;

    const trimmed = username.trim();

    let resp = await check(updateDetails(token, trimmed, password));

    if (resp.error) {
      setState({ ...state, error: resp.message, success: '' });
      return
    }

    setToken(resp.token);
    setState({ password: '', error: '', success: 'Username updated' });
  };

  const downloadData = async () => {
    console.log(await DownloadData(context.token));
  };

  return (
    <>
      <Tab.Pane eventKey="profile">
        <h5>Profile</h5>
        <hr />
        {state.success !== '' && 
          <Alert
            variant='success'
          >
            {state.success}
          </Alert>
        }
        <Form onSubmit={updateProfile}>
          <Row>
            <Form.Group as={Col}>
              <Form.Label>Username</Form.Label>
              <Form.Control
                className='mb-3'
                type='username'
                name='username'
                placeholder='Username'
                onChange={updateState}
                value={state.username}
                required
              />
              <Form.Label>Password</Form.Label>
              <Form.Control
                className='mb-3'
                type='password'
                name='password'
                placeholder='Password'
                onChange={updateState}
                value={state.password}
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
            <Col lg='3' md='5' sm='12' xs='12'>
              <Button
                className='w-100'
                type='submit'
              >
                Save Changes
              </Button>
            </Col>
          </Row>
        </Form>
        <p></p>
        <h5>Export Data</h5>
        <hr />
        <Row className='justify-content-end'>
          <Col>
            Export user data into a JSON file
          </Col>
          <Col lg='2' md='5' sm='12' xs='12'>
            <Button
              className='w-100'
              variant='primary'
              onClick={downloadData}
            >
            Export JSON
          </Button>
          </Col>
        </Row>
      </Tab.Pane>
    </>
  )
}

export default Profile;
