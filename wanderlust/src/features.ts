export async function handleCount(count: any, setCount: any) {
  const response = await fetch("http://localhost:3001/api/v1/count/handle", {
    method: "POST",
    body: JSON.stringify({ count: count }),
    headers: {
      "Content-Type": "application/json",
    },
  }).then((res) => {
    return res.json();
  });
  if (response && !response.error) {
    setCount(response.data.count);
  }
}

export async function handleForm(
  company: any,
  contact: any,
  country: any,
  setCompanys: any
) {
  const response = await fetch("http://localhost:3001/api/v1/form/handle", {
    method: "POST",
    body: JSON.stringify({
      company: company,
      contact: contact,
      country: country,
    }),
    headers: {
      "Content-Type": "application/json",
    },
  }).then((res) => {
    return res.json();
  });
  if (response && !response.error) {
    setCompanys(response.data);
  }
}

export async function refreshUsers(users: any, setUsers: any) {
  const response = await fetch("http://localhost:3001/api/v1/users/all", {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  }).then((res) => {
    return res.json();
  });
  if (response && !response.error) {
    setUsers(response.data);
  }
}

export async function createUser(userValues: {
  firstname: any;
  lastname: any;
  username: any;
  email: any;
  description: any;
}) {
  const response = await fetch("http://localhost:3001/api/v1/users/create", {
    method: "POST",
    body: JSON.stringify({
      firstname: userValues.firstname,
      lastname: userValues.lastname,
      username: userValues.username,
      email: userValues.email,
      description: userValues.description,
    }),
    headers: {
      "Content-Type": "application/json",
    },
  }).then((res) => {
    return res.json();
  });
}
