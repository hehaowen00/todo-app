import React, { useContext, useState } from 'react'
import { Alert, Button, Container, Form, FormControl, Row, Col } from 'react-bootstrap'

import Topbar from '../util/Topbar'

import { AuthContext } from '../../context/auth';
import { AddList } from '../../api/todo';
import { useNavigate } from 'react-router-dom';

function NewList() {
  const {context} = useContext(AuthContext);
  const nav = useNavigate();

  const [error, setError] = useState(null);
  const [name, setName] = useState('');

  const handleInput = (e) => {
    setName(e.target.value);
  };

  const addNewList = async (e) => {
    e.preventDefault();
    setError('');

    if (name && name.trim() !== '') {
      let resp = await AddList(context.token, name.trim());
      if (resp.error) {
        setError('Failed to add list');
        return
      }

      nav('/');
    }
  };

  const cancel = (e) => {
    window.history.back();
  };

  return (
    <>
      <Topbar />
      <Container>
        <p></p>
        {error && 
        <>
          <Alert variant='danger'>
            {error}
          </Alert>
          <p></p>
        </>
}
        <Form onSubmit={addNewList}>
          <Row className='d-flex'>
            <Col md='12'>
            <FormControl
              type='input'
              placeholder='New List'
              className='me-2 add-margin'
              aria-label='Name'
              onChange={handleInput}
              value={name}
            />
            </Col>
            </Row>

        <Row className='justify-content-end'> 
            <Col md='2'>
              <Button
                className='form-control'
                variant='outline-primary'
                type='submit'
              >
                Add List
              </Button>
            </Col>
            <Col md='2'>
              <Button
                className='form-control'
                variant='outline-danger'
                onClick={cancel}
              >
                Cancel
              </Button>
            </Col>
          </Row>
        </Form>
      </Container>
    </>
  )
}

export default NewList;