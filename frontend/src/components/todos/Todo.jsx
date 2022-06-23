import React, { useContext, useEffect, useState } from 'react';
import { Alert, Button, Container, Form, FormControl, Row, Col, FormGroup } from 'react-bootstrap';
import { useParams } from 'react-router-dom';

import Topbar from '../util/Topbar';

import { AuthContext } from '../../context/Auth';
import { GetTodo, UpdateTodo } from '../../api/Todo';
import { setTitle } from '../../util/Util';

function Todo() {
  const { context, check } = useContext(AuthContext);
  const { id } = useParams();

  const [state, setState] = useState({
    error: '',
    id: parseInt(id),
    desc: '',
    status: false,
  });
  
  const cancel = () => {
    window.history.back();
  };

  const updateState = (e) => {
    const { name, value } = e.target;
    setState({ ...state, [name]: value });
  };

  const toggleStatus = (e) => {
    setState({ ...state, status: !state.status });
  };

  const getItem = async () => {
    let resp = await check(GetTodo(context.token, parseInt(id)));
    if (resp.error) {
      cancel();
      return;
    }

    setState(resp.item);
  };

  const updateItem = async (e) => {
    e.preventDefault();
    
    setState({ ...state, error: '' });
    let resp = await check(UpdateTodo(context.token, state));
    if (resp.error) {
      setState({ ...state, error: resp.message });
      return;
    }

    window.history.back();
  };

  useEffect(() => {
    setTitle('Edit Item');
    getItem();
  }, []);

  return (
    <>
      <Topbar />
      <Container>
        <p></p>
        {state.error && 
        <>
          <Alert variant='danger'>
            {state.error}
          </Alert>
          <p></p>
        </>}
        <Form onSubmit={updateItem}>
          <Row>
            <Col>
              <FormGroup className='mb-3'>
                <FormControl
                  type='input'
                  className='me-2'
                  name='desc'
                  placeholder='Todo'
                  onChange={updateState}
                  required
                  value={state.desc}
                />
              </FormGroup>
            </Col>
          </Row>
          <Row className='justify-content-center'>
            <Col md='10'>
              <FormGroup>
                <Form.Check
                  className='ml-5'
                  type='checkbox'
                  label='Completed'
                  checked={state.status}
                  onChange={toggleStatus}
                />
              </FormGroup>
            </Col>
            <Col md='2'>
              <FormGroup>
                <Button
                  className='w-100'
                  variant='primary'
                  type='submit'
                >
                  Save
                </Button>
              </FormGroup>
            </Col>
          </Row>
        </Form>
      </Container>
    </> 
  )
}

export default Todo;
