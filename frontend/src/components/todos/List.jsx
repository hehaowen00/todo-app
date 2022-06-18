import React, { useCallback, useContext, useEffect, useState } from 'react';
import { Alert, Button, Container, Form, FormControl, ListGroup, Row, Col } from 'react-bootstrap'
import { useParams } from 'react-router-dom';

import TodoItem from './TodoItem';
import Topbar from '../util/Topbar';

import { AuthContext } from '../../context/auth';
import { ListContext } from '../../context/list';
import { AddTodo, DeleteTodo, GetTodos, UpdateList, UpdateTodo } from '../../api/todo';

function List() {
  const { context } = useContext(AuthContext);
  const { getList } = useContext(ListContext);
  const { id } = useParams();

  const list = getList(parseInt(id));
  const [state, setState] = useState({
     desc: '',
     error: '',
     filter: '1',
     name: list.name,
     todos: [],
  });

  const filter = (e) => {
    switch (state.filter) {
      case '2':
        return !e.status;
      case '3':
        return e.status === true;
      default:
        return true;
    }
  }

  let timer;

  const keydown = () => {
    clearTimeout(timer);
  };

  const keyup = () => {
    clearTimeout(timer);
    timer = setTimeout(updateList, 1000);
  };

  const updateState = (e) => {
    const { name, value } = e.target;
    setState({ ...state, [name]: value });
  };

  const updateList = async () => {
    let resp = await UpdateList(context.token, { ...list, name: state.name });
    if (resp.error) {
      setState({ ...state, error: 'Unable to update list name'});
      return;
    }
  };

  const addTodo = async (e) => {
    e.preventDefault();
    if (state.desc === '') {
      return;
    }

    let resp = await AddTodo(context.token, list.id, state.desc);
    if (resp.error) {
      setState({ ...state, error: resp.message });
      return;
    }
    setState({
      ...state,
      desc: '',
      todos: [...state.todos, resp.item ],
    });
  };

  const updateTodo = async (todo) => {
    let resp = await UpdateTodo(context.token, todo);
    if (resp.error) {
      return;
    }

    let todos = [...state.todos];
    let idx = todos.findIndex(e => e.id === todo.id);
    todos[idx] = resp.item;

    setState({ ...state, todos, });
  };

  const deleteTodo = async (todo) => {
    let resp = await DeleteTodo(context.token, todo);

    if (resp.error) {
      return;
    }

    let todos = [...state.todos];
    let idx = todos.findIndex(e => e.id === todo.id);
    todos.splice(idx, 1);

    setState({ ...state, todos, });
  };

  const getTodos = useCallback(async () => {
    let resp = await GetTodos(context.token, list.id);
    if (resp.error) {
      return;
    }
    setState({ ...state, todos: resp.todos});
  }, []);

  useEffect(() => {
    getTodos();
  }, [getTodos]);

  return (
    <>
      <Topbar />
      <Container>
        <p></p>
        <Form>
          <Row>
          <Col md='10'>
            <FormControl
              type='input'
              className='me-2'
              name='name'
              placeholder='Name'
              onChange={updateState}
              onKeyDown={keydown}
              onKeyUp={keyup}
              value={state.name}
            />
          </Col>
          <Col>
            <Form.Select name='filter' size='xs' onChange={updateState}>
              <option value='1'>All</option>
              <option value='2'>In Progress</option>
              <option value='3'>Done</option>
            </Form.Select>
          </Col>
          </Row>
        </Form>
        <p></p>
        <Form onSubmit={addTodo}>
          <Row>
            <Col md='11'>
              <FormControl
                type='input'
                className='me-2'
                name='desc'
                onChange={updateState}
                placeholder='New Todo'
                required={true}
                value={state.desc}
              />
            </Col>
            <Col>
              <Button
              className='w-100'
                type='submit'
                onClick={addTodo}
                variant='primary'
              >
                Add
              </Button>
            </Col>
          </Row>
        </Form>
        <p></p>
        {state.error !== '' &&
          <Alert variant='danger'>
            {state.error}
          </Alert>
         }
        <ListGroup>
          {state.todos.filter(item => filter(item)).map((todo, idx) =>
            <TodoItem
              key={idx}
              item={todo}
              remove={deleteTodo}
              update={updateTodo} />)}
        </ListGroup>
        <p></p>
      </Container>
    </>
  )
}

export default List;