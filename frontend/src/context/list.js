import React, { useState } from 'react';

export const ListContext = React.createContext();

export function ListProvider({ children }) {
  const [context, setState] = useState({ lists: [], });

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
    <ListContext.Provider value={{ context, addLists, getList, removeList }}>
      {children}
    </ListContext.Provider>
  )
}
