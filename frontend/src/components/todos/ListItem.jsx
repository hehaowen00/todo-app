import React, { useContext } from 'react';
import { ListGroupItem } from 'react-bootstrap';
import { Link } from 'react-router-dom';

import { AuthContext } from '../../context/auth';
import { ListContext } from '../../context/list';
import { DeleteList } from '../../api/todo';

function ListItem({ id, name }) {
  const { context } = useContext(AuthContext);
  const { removeList } = useContext(ListContext);
  const link = `/list/${id}`;

  const deleteList = async (e) => {
    let resp = await DeleteList(context.token, {id, name});
    if (resp.error) {
    }
    removeList({ id, name });
  };

  return (
    <ListGroupItem
      className='d-flex justify-content-between align-items-start list-item'
    >
      <div className='w-100'>
      <Link className='plain blue' to={link}>
        <div className='item-text no-select pointer'>
          {name}
        </div>
      </Link>
      </div>
      <button
        className='link red'
        onClick={deleteList}
      >
        <b>X</b>
      </button>
    </ListGroupItem>
  )
}

export default ListItem;