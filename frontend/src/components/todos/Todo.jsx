import React, { useContext, useEffect, useState } from 'react';
import { Button, Container, Form, FormControl, Row, Col, FormGroup } from 'react-bootstrap';
import { useParams } from 'react-router-dom';
import { GetTodo, UpdateTodo } from '../../api/todo';
import { AuthContext } from '../../context/auth';
import Topbar from '../util/Topbar';

function Todo() {
  const { context } = useContext(AuthContext);
  const { id } = useParams();
  const [state, setState] = useState({
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
    let resp = await GetTodo(context.token, parseInt(id));
    if (resp.error) {
      return;
    }

    setState(resp.item);
  };

  const updateItem = async (e) => {
    e.preventDefault();
    let resp = await UpdateTodo(context.token, state);
    if (resp.error) {
      return;
    }

    window.history.back();
  };

  useEffect(() => {
    getItem();
  }, []);

  return (
    <>
      <Topbar />
      <Container>
        <p></p>
        <Form onSubmit={updateItem}>
          <Row>
            <Col>
              <FormGroup className='mb-3'>
                <FormControl
                  type='input'
                  className='me-2'
                  name='desc'
                  placeholder='Todo'
                  required
                  onChange={updateState}
                  value={state.desc}
                />
              </FormGroup>
            </Col>
          </Row>
          <Row className='justify-content-center'>
            <Col md='8'>
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
              variant='outline-primary'
              type='submit'
            >
              Save
            </Button>
            </FormGroup>
            </Col>
            <Col md='2'>
            <FormGroup>
              <Button
                className='w-100'
                variant='outline-danger'
                onClick={cancel}
              >
                Cancel
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