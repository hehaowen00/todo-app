import React, { useContext, useEffect, useState } from 'react';
import { Alert, Container, Form, FormControl, ListGroup, Row, Col } from 'react-bootstrap';

import ListItem from './todos/ListItem';
import Topbar from './util/Topbar';

import { ListContext } from '../context/List';

function Home() {
  const { context, getLists } = useContext(ListContext);
  const { error, lists} = context;

  const [search, setSearch] = useState('');

  const match = (s) => {
    return s.toLowerCase().includes(search.toLowerCase());
  };

  useEffect(() => {
    getLists()
  }, [])

  return (
      <>
        <Topbar />
        <Container>
          <p></p>
          <Form>
            <Row>
              <Col className='add-margin' md='12'>
                <FormControl
                  type='search'
                  placeholder='Find Todo List'
                  className='me-3'
                  aria-label='search'
                  value={search}
                  onChange={e => setSearch(e.target.value)}
                />
              </Col>
            </Row>
          </Form>
          <p></p>
          {error !== '' && 
            <>
              <Alert variant='danger'>
                {error}
              </Alert>
            </>
          }
          <ListGroup as='ol'>
            {lists.filter(e => match(e.name)).map(list=> 
              <ListItem key={list.id} id={list.id} name={list.name} count={0} />
            )}
          </ListGroup> 
          <p></p> 
        </Container>
      </>
  )
}

export default Home;