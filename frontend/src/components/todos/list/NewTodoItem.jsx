import React, { useState } from 'react';
import { Button, Col, Form, FormControl, Row} from 'react-bootstrap';

import { AddTodo, } from '../../../api/Todo';

function NewTodoItem({ context, check, id, todoState, setError }) {
  const [desc, setDesc] = useState('');
  const [todos, setTodos] = todoState;

  const addTodo = async (e) => {
    e.preventDefault();

    const desc_ = desc.trim();

    if (desc_ === '') {
      return;
    }

    let resp = await check(AddTodo(context.token, id, desc_));
    if (resp.error) {
      setError(resp.message);
      return;
    }

    setDesc('');
    setTodos([...todos, resp.item ]);
  };

  return (
    <Form onSubmit={addTodo}>
      <Row>
        <Col md='10'>
          <FormControl
            type='input'
            className='me-2'
            autoComplete='off'
            name='desc'
            onChange={e => setDesc(e.target.value)}
            placeholder='Item'
            required={true}
            value={desc}
          />
        </Col>
        <Col>
          <Button
            className='w-100'
            type='submit'
            onClick={addTodo}
            variant='primary'
            disabled={desc === ''}
          >
            Add Item
          </Button>
        </Col>
      </Row>
    </Form>
  )
}

export default NewTodoItem;
