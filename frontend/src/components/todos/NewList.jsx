import React, { useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom';
import { Alert, Button, Container, Form, FormControl, Row, Col } from 'react-bootstrap'

import Topbar from '../util/Topbar'

import { AuthContext } from '../../context/Auth';
import { AddList } from '../../api/List';
import { setTitle } from '../../util/Util';

function NewList() {
  const { context, check } = useContext(AuthContext);
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
      let resp = await check(AddList(context.token, name.trim()));
      if (resp.error) {
        setError(resp.message);
        return
      }

      nav('/');
    }
  };

  useEffect(() => {
    setTitle('New List');
  });

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
        </>}
        <Form onSubmit={addNewList}>
          <Row className='d-flex'>
            <Col md='12'>
            <FormControl
              type='input'
              placeholder='New List'
              className='me-2 add-margin'
              aria-label='Name'
              onChange={handleInput}
              required
              value={name}
            />
            </Col>
          </Row>
          <Row className='justify-content-end'> 
            <Col md='2'>
              <Button
                className='form-control'
                variant='primary'
                type='submit'
              >
                Add List
              </Button>
            </Col>
          </Row>
        </Form>
      </Container>
    </>
  )
}

export default NewList;
