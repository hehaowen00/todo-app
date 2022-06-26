import React, { useCallback, useEffect } from 'react';
import { ListGroup } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';

import TodoItem from '../TodoItem';

import { GetTodos, UpdateTodo, DeleteTodo } from '../../../api/Todo';

function ListRender({ context, check, id, todoState, filter }) {
  const [todos, setTodos] = todoState;
  const nav = useNavigate();

  const updateTodo = async (todo) => {
    let resp = await check(UpdateTodo(context.token, todo));

    if (resp.unauthorized === true) {
      nav('/');
      return;
    }

    if (resp.error) {
      return;
    }

    let todos_ = [...todos];
    let idx = todos_.findIndex(e => e.id === todo.id);
    todos_[idx] = resp.item;

    setTodos(todos_);
  };

  const deleteTodo = async (todo) => {
    let resp = await check(DeleteTodo(context.token, todo));

    if (resp.unauthorized === true) {
      nav('/');
      return;
    }

    if (resp.error) {
      return;
    }

    let todos_ = [...todos];
    let idx = todos_.findIndex(e => e.id === todo.id);
    todos_.splice(idx, 1);

    setTodos(todos_);
  };

  const getTodos = useCallback(async () => {
    let resp = await check(GetTodos(context.token, id));

    if (resp.unauthorized === true) {
      nav('/');
      return;
    }

    if (resp.error) {
      return;
    }

    setTodos(resp.todos)
  }, []);

  useEffect(() => {
    getTodos();
  }, []);

  return (
    <ListGroup>
      {todos.filter(item => filter(item)).map((todo, idx) =>
        <TodoItem
          key={idx}
          item={todo}
          remove={deleteTodo}
          update={updateTodo}
        />
      )}
    </ListGroup>
  )
}

export default ListRender;