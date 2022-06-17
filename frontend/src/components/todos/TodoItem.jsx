import React from 'react';
import { Form, ListGroupItem } from 'react-bootstrap';
import { useNavigate } from 'react-router-dom';

function TodoItem({ item, remove, update }) {
  const nav = useNavigate();

  const style = { textDecoration: item.status ? 'line-through': ''};

  const editTodo = (e) => {
    const link = `/item/${item.id}`
    nav(link);
  };

  const updateStatus = async (e) => {
    let newItem = { ...item, status: !item.status };
    update(newItem);
  };

  const deleteTodo = async (e) => {
    remove(item);
  };

  return (
    <ListGroupItem
      className='d-flex justify-content-between align-items-start'
    >
      <Form>
        <Form.Check
          className='ml-5'
          type='checkbox'
          checked={item.status}
          onChange={updateStatus}
        />
      </Form>
      <div className='w-100' style={style} onClick={editTodo}>
        <div className='item-text no-select pointer'>
          <span>{item.desc}</span>
        </div>
      </div>
      <button
        className='link red'
        onClick={deleteTodo}
      >
        <b>X</b> 
      </button>
    </ListGroupItem>
  )
}

export default TodoItem;