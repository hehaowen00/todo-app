import React from 'react';
import { Container, Row, Col } from 'react-bootstrap';
import { Link } from 'react-router-dom';

import Topbar from './Topbar';

function NotFound() {
  return (
    <>
      <Topbar />
      <Container>
        <p></p>
        <Row className='justify-content-center'>
          <Col className='text-center'>
            <h2>Page Not Found</h2>
            <p>Click <Link to='/'>Here</Link> To Return</p>
          </Col>
        </Row>
      </Container>
    </>
  )
}

export default NotFound;