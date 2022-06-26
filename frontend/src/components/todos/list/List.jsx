import React, { useContext, useState } from 'react';
import { Alert, Container } from 'react-bootstrap'
import { useParams } from 'react-router-dom';

import ListHeader from './ListHeader';
import ListRender from './ListRender';
import NewTodoItem from './NewTodoItem';
import Topbar from '../../util/Topbar';

import { AuthContext } from '../../../context/Auth';

function List() {
  const { context, check } = useContext(AuthContext);
  const { id } = useParams();
  const ID = parseInt(id);

  const [error, setError] = useState('');
  const [opt, setOpt] = useState('1');
  const todosHook = useState([]);

  const filter = (el) => {
    switch (opt) {
      case '2':
        return !el.status;
      case '3':
        return el.status === true;
      default:
        return true;
    }
  };

  return (
    <>
      <Topbar />
      <Container>
        <p></p>
        <ListHeader
          context={context}
          check={check}
          id={ID}
          setError={setError}
          setOpt={setOpt}
        />
        <p></p>
        <NewTodoItem
          context={context}
          check={check}
          id={ID}
          todoState={todosHook}
          setError={setError}
        />
        <p></p>
        {error !== '' &&
          <Alert variant='danger'>
            {error}
          </Alert>
        }
        <ListRender
          context={context}
          check={check}
          id={ID}
          todoState={todosHook}
          filter={filter}
        />
        <p></p>
      </Container>
    </>
  )
}

export default List;
