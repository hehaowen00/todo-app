import React, { useCallback, useEffect, useState } from 'react';
import { GetLists } from '../api/List';

export const ListContext = React.createContext();

export function ListProvider({ children }) {
  const [context, setState] = useState({ lists: [], error: '', });

  const getLists = useCallback(async () => {
    setError('');
    let token = localStorage.getItem('token');
    let resp = await GetLists(token);

    if (resp.error === true) {
      setError('Failed to get todo lists');
      return resp;
    }

    setState({ ...context, lists: resp.lists });
  }, []);

  const setError = (error) => {
    setState({ ...context, error })
  }

  useEffect(() => {
    getLists();
  }, []);

  const addLists = (lists) => {
    setState({ ...context, lists });
  };

  const getList = (id) => {
    const { lists } = context;
    const idx = lists.findIndex(e => e.id === id);
    return lists[idx];
  };

  const removeList = (list) => {
    const { lists } = context;
    const idx = lists.findIndex(e => e.id === list.id);
    lists.splice(idx, 1);
    setState({ ...context, lists });
  };

  return (
    <ListContext.Provider value={{ context, addLists, getLists, getList, removeList }}>
      {children}
    </ListContext.Provider>
  )
}
