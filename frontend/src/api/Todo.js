const LIST_ENDPOINT = '/api/lists';
const TODO_ENDPOINT = '/api/todos';
const ITEM_ENDPOINT = '/api/todo';

export async function GetLists(token) {
  try {
    let resp = await fetch(LIST_ENDPOINT, {
      method: 'GET',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (resp.ok) {
      let lists = await resp.json();
      return { error: false, lists};
    }

    return { error: true, message: 'Failed to get lists' };
  } catch (error) {
    return {
       error: true,
       message: 'Failed to connect to server'
    };
  }
}

export async function GetList(token, id) {
  if (id === undefined) {
    return {error: true }
  }
  try {
    let resp = await fetch(LIST_ENDPOINT + `/${id}`, {
      method: 'GET',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (resp.ok) {
      let list = await resp.json();
      return { error: false, list};
    }

    return { error: true, message: 'Failed to get lists' };
  } catch (error) {
    return {
       error: true,
       message: 'Failed to connect to server'
    };
  }
}

export async function AddList(token, name) {
  try {
    let resp = await fetch(LIST_ENDPOINT, {
      method: 'POST',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        name,
      }),
    });

    if (resp.ok) {
      return { error: false };
    }

    return { error: true, message: 'Failed to add list' };
  } catch (error) {
    return {
       error: true,
       message: 'Failed to connect to server'
    };
  }
}

export async function UpdateList(token, list) {
  try {
    let resp = await fetch(LIST_ENDPOINT, {
      method: 'PUT',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify(list),
    });

    if (resp.ok) {
      let list = await resp.json();
      return { error: false, list };
    }

    return { error: true, message: 'Failed to update list' };
  } catch (error) {
    return {
       error: true,
       message: 'Failed to connect to server'
    };
  }
}

export async function DeleteList(token, list) {
  try {
    let resp = await fetch(LIST_ENDPOINT, {
      method: 'DELETE',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify(list)
    });

    if (resp.ok) {
      let lists = await resp.json();
      return { error: false, lists};
    }

    return { error: true, message: 'Failed to delete list' };
  } catch (error) {
    return {
       error: true,
       message: 'Failed to connect to server'
    };
  }
}

export async function GetTodos(token, listId) {
  try {
    let resp = await fetch(TODO_ENDPOINT + `/${listId}`, {
      method: 'GET',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (resp.ok) {
      let todos = await resp.json();
      return { error: false, todos };
    }

    return { error: true, message: 'Failed to get todos' };
  } catch (error) {
    return {
       error: true,
       message: 'Failed to connect to server'
    };
  }
}

export async function GetTodo(token, itemId) {
  try {
    let resp = await fetch(ITEM_ENDPOINT + `/${itemId}`, {
      method: 'GET',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (resp.ok) {
      let item = await resp.json();
      return { error: false, item };
    }

    return { error: true, message: 'Failed to get todo item' };
  } catch (error) {
    return {
       error: true,
       message: 'Failed to connect to server'
    };
  }
}

export async function AddTodo(token, listId, desc) {
  try {
    let resp = await fetch(TODO_ENDPOINT, {
      method: 'POST',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        'list_id': listId,
        'desc': desc,
      })
    });

    if (resp.ok) {
      let item = await resp.json();
      return { error: false, item };
    }

    return { error: true, message: 'Failed to add item' };
  } catch (error) {
    return {
       error: true,
       message: 'Failed to connect to server'
    };
  }
}

export async function UpdateTodo(token, todo) {
  try {
    let resp = await fetch(TODO_ENDPOINT, {
      method: 'PUT',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify(todo),
    });

    if (resp.ok) {
      let item = await resp.json();
      return { error: false, item };
    }

    return { error: true, message: 'Failed to update item' };
  } catch (error) {
    return {
       error: true,
       message: 'Failed to connect to server'
    };
  }
}

export async function DeleteTodo(token, todo) {
  try {
    let resp = await fetch(TODO_ENDPOINT, {
      method: 'DELETE',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify(todo),
    });

    if (resp.ok) {
      return { error: false };
    }

    return { error: true, message: 'Failed to delete item' };
  } catch (error) {
    return {
       error: true,
       message: 'Failed to connect to server'
    };
  }
}