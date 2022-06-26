const LIST_ENDPOINT = '/api/lists';

export async function GetLists(token) {
  try {
    let resp = await fetch(LIST_ENDPOINT, {
      method: 'GET',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (resp.status === 401) {
      return { error: true, unauthorized: true };
    }

    if (resp.ok) {
      let lists = await resp.json();
      return { error: false, lists};
    }

    return { error: true, message: await resp.json() };
  } catch (error) {
    return {
       error: true,
       unauthorized: true,
       message: 'Failed to connect to server'
    };
  }
}

export async function GetList(token, id) {
  if (id === undefined) {
    return {error: true }
  }
  try {
    let resp = await fetch(`${LIST_ENDPOINT}/${id}`, {
      method: 'GET',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (resp.status === 401) {
      return { error: true, unauthorized: true };
    }

    if (resp.ok) {
      let list = await resp.json();
      return { error: false, list};
    }

    return { error: true, message: await resp.json() };
  } catch (error) {
    return {
       error: true,
       unauthorized: true,
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

    if (resp.status === 401) {
      return { error: true, unauthorized: true };
    }

    if (resp.ok) {
      return { error: false };
    }

    return { error: true, message: await resp.json() };
  } catch (error) {
    return {
       error: true,
       unauthorized: true,
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

    if (resp.status === 401) {
      return { error: true, unauthorized: true };
    }

    if (resp.ok) {
      let list = await resp.json();
      return { error: false, list };
    }

    return { error: true, message: await resp.json() };
  } catch (error) {
    return {
       error: true,
       unauthorized: true,
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

    if (resp.status === 401) {
      return { error: true, unauthorized: true };
    }

    if (resp.ok) {
      let lists = await resp.json();
      return { error: false, lists};
    }

    return { error: true, message: await resp.json() };
  } catch (error) {
    return {
       error: true,
       unauthorized: true,
       message: 'Failed to connect to server'
    };
  }
}
