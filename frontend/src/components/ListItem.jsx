import React, { useContext, useState } from 'react';
import { Button, ListGroupItem, Modal } from 'react-bootstrap';
import { Link } from 'react-router-dom';

import { AuthContext } from '../context/Auth';
import { ListContext } from '../context/List';
import { DeleteList } from '../api/List';

function ListItem({ id, name }) {
  const { context, check } = useContext(AuthContext);
  const { removeList } = useContext(ListContext);

  const [show, setShow] = useState(false);

  const link = `/list/${id}`;

  const deleteList = async (e) => {
    let resp = await check(DeleteList(context.token, {id, name}));
    if (resp.error) {
      return;
    }
    removeList({ id, name });
  };

  const toggleModal = (e) => {
    setShow(!show);
  };

  return (
    <>
    <Modal show={show} centered>
      <Modal.Body>
        Delete "{name}" ?
      </Modal.Body>
      <Modal.Footer>
        <Button
         onClick={deleteList}
         variant='danger'
         size='sm'
        >
          Delete
        </Button>
        <Button
         onClick={toggleModal}
         variant='primary'
         size='sm'
        >
          Cancel
        </Button>
      </Modal.Footer>
    </Modal>
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
        onClick={toggleModal}
      >
        <b>X</b>
      </button>
    </ListGroupItem>
    </>
  )
}

export default ListItem;