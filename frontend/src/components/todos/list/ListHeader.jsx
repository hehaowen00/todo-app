import React, { useCallback, useEffect, useState } from 'react';
import { Col, Form, FormControl, Row } from 'react-bootstrap';
import { Timer } from '../../../util/Util';

import { GetList, UpdateList } from '../../../api/List';
import { useNavigate } from 'react-router-dom';

function ListHeader({ context, check, id, setError, setOpt }) {
  const [name, setName] = useState('');
  const [list, setList] = useState({name: ''});
  const nav = useNavigate();

  const nameTimer = new Timer();

  const keydown = () => {
    nameTimer.keydown();
  };

  const keyup = () => {
    nameTimer.keyup(updateList, 500);
  };

  const updateList = async () => {
    const name = name.trim();
    let resp = await check(UpdateList(context.token, { ...list, name }));

    if (resp.unauthorized === true) {
      nav('/');
      return;
    }

    if (resp.error) {
      setError(resp.message);
      return;
    }
  };

  const getList = useCallback(async () => {
    let resp = await check(GetList(context.token, id));
    if (resp.error) {
      nav('/');
      return;
    }

    const { list } = resp;
    setList(list);
    setName(list.name);
  }, []);

  useEffect(() => {
    getList();
  }, []);

  return (
    <Form>
      <Row>
      <Col md='10'>
        <FormControl
          type='input'
          className='me-2'
          name='name'
          placeholder='Name'
          onChange={setName}
          onKeyDown={keydown}
          onKeyUp={keyup}
          value={list.name}
        />
      </Col>
      <Col>
        <Form.Select name='filter' size='xs' onChange={e => setOpt(e.target.value)}>
          <option value='1'>All</option>
          <option value='2'>In Progress</option>
          <option value='3'>Done</option>
        </Form.Select>
      </Col>
      </Row>
    </Form>
  )
}

export default ListHeader;