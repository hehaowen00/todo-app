import React, { useEffect } from 'react';
import { Container, Row, Col, Tab, Nav } from 'react-bootstrap';

import DeleteAccount from './settings/DeleteAccount';
import Security from './settings/Security';
import Profile from './settings/Profile';
import Topbar from './util/Topbar';

import { setTitle } from '../util/Util';

function Settings() {
  useEffect(() => {
    setTitle('Settings');
  });

  return (
    <>
      <Topbar />
      <Container>
        <p></p>
        <Row className='settings-title'>
          <Col md={6}>
            <h4>Settings</h4>
          </Col>
        </Row>
        <Tab.Container defaultActiveKey='profile'>
          <Row>
            <Col className='settings-header' sm='3'>
              <Nav variant="pills" className="flex-column">
              <Nav.Item>
                <Nav.Link eventKey="profile">Profile</Nav.Link>
              </Nav.Item>
              <Nav.Item>
                <Nav.Link eventKey="security">Security</Nav.Link>
              </Nav.Item>
              <div className="border"></div>
              <Nav.Item className='delete-pill'>
                <Nav.Link className='red' eventKey="delete">Delete Account</Nav.Link>
              </Nav.Item>
              </Nav>
            </Col>
            <Col sm={9}>
              <Tab.Content>
                <Profile />
                <Security />
                <DeleteAccount />
              </Tab.Content>
            </Col>
          </Row>
        </Tab.Container>
      </Container>
    </>
  )
}

export default Settings;
