import React, { useState } from 'react';
import { GetTodos } from '../api/Todo';

export const TodosContext = React.createContext();

export function TodosProvider({ children }) {
  const [context, setContext] = useState({
    listId: null, 
    todos: [],
  });

  const useTodos = async (listId) => {
    let token = localStorage.getItem('token');
    let resp = await GetTodos(token);
    if (resp.error === true) {
      return;
    }

    setContext({ listId, todos: resp.todos });
  };

  return (
    <TodosContext.Provider value={{context, useTodos}}>
      {children}
    </TodosContext.Provider>
  )
}