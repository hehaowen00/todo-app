import React, { useContext } from 'react';
import { Container, Nav, Navbar } from 'react-bootstrap'
import { LinkContainer } from 'react-router-bootstrap';
import { Navigate} from 'react-router-dom';

import { AuthContext } from '../../context/auth';

function Topbar() {
  const { context }= useContext(AuthContext);

  if (context.allowed) {
    if (context.payload) {
      const { username } = context.payload;
      return <SignedIn name={username} />
    }
    return <Navigate to='/' />
  }

  return <Default />
}

function Default() {
  return (
    <Navbar bg='dark' variant='dark' expand='lg'>
      <Container>
        <LinkContainer to='/'>
          <Navbar.Brand className='white'>
            Todo App
          </Navbar.Brand>
        </LinkContainer>
        <Navbar.Toggle aria-controls='navbar'/>
        <Navbar.Collapse id='navbar'>
          <Nav className='me-auto my-lg-0'>
          </Nav>
          <Nav>
            <LinkContainer to='/'>
              <Nav.Link className='white-link'>Login</Nav.Link>
            </LinkContainer>
            <LinkContainer to='/register'>
              <Nav.Link className='white-link'>Register</Nav.Link>
            </LinkContainer>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  )
}

function SignedIn({ name }) {
  return (
    <Navbar bg='dark' variant='dark' expand='lg' sticky='top'>
      <Container>
        <LinkContainer to='/'>
          <Navbar.Brand className='white'>
            Todo App
          </Navbar.Brand>
        </LinkContainer>
        <Navbar.Toggle aria-controls='navbar'/>
        <Navbar.Collapse id='navbar'>
          <Nav className='me-auto'>
            <LinkContainer to='/home'>
              <Nav.Link>My Lists</Nav.Link>
            </LinkContainer>
            <LinkContainer to='/newList'>
              <Nav.Link>Add List</Nav.Link>
            </LinkContainer>
          </Nav>
          <Nav>
          <Navbar.Text className='white'>
            Signed in as: <span style={{textDecoration: 'underline'}}>{name}</span>
          </Navbar.Text>
          <LinkContainer to='/logout'>
            <Nav.Link>Sign Out</Nav.Link>
          </LinkContainer>
          </Nav>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  )
}

export default Topbar;