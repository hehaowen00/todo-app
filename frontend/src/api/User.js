const BASE_URL = '/api/user'

export async function updateDetails(token, username, password) {
  try {
    let resp = await fetch(`${BASE_URL}/profile`, {
      method: 'PUT',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        new_username: username,
        password,
      }),
    });

    console.log(resp);

    if (resp.status === 401) {
      return {
        error: true,
        unauthorized: true,
      };
    }

    if (resp.ok) {
      let token = await resp.json();
      return { error: false, token };
    }

    let message = await resp.json();
    return { error: true, message };
  } catch (error) {
    return { error: true, message: "Failed to connect to server" };
  }
}

export async function updatePassword(token, password, new_password) {
  try {
    let resp = await fetch(`${BASE_URL}/password`, {
      method: 'PUT',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        password,
        new_password
      }),
    });

    let message = await resp.json();

    if (resp.ok) {
      return { error: false, message };
    }

    return { error: true, message };
  } catch (error) {
    return { error: true, message: 'Failed to connect to server' };
  } 
}

export async function deleteAccount(token, password) {
  try {
    let resp = await fetch(`${BASE_URL}`, {
      method: 'DELETE',
      mode: 'cors',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        password
      }),
    });

    let message = await resp.json();

    if (resp.ok) {
      return {
        error: false,
        message
      };
    }

    if (resp.status === 401) {
      return {
        error: true,
        unauthorized: true,
      };
    }

    return { error: true, message };
  } catch (error) {
    return { error: true, message: 'Failed to connect to server' };
  }
}

export async function DownloadData(token) {
  try {
    let resp = await fetch(`${BASE_URL}/data`, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (resp.status === 401) {
      return { error: true, unauthorized: true };
    }

    if (!resp.ok) {
      let message = await resp.json();
      return { error: true, message };
    }

    let data = await resp.json();
    let blob = new Blob([JSON.stringify(data, null, 2)], {
      type: 'application/json',
    });

    var a = document.createElement('a');
    a.href = window.URL.createObjectURL(blob);
    a.download = 'data.json';
    a.click();

    return { error: true };
  } catch (error) {
    return { error: true, message: 'Failed to connect to server' };
  }
}
