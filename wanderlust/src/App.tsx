import { useEffect, useState } from "react";
import { createUser, handleCount, handleForm, refreshUsers } from "./features";

interface Users {
  firstname: string;
  lastname: string;
  username: string;
  email: string;
  description: string;
}

export default function App() {
  const [count, setCount] = useState(0);
  const [companys, setCompanys] = useState([
    {
      company: "Alfreds Futterkiste",
      contact: "Maria Anders",
      country: "Germany",
    },
  ]);
  const [formCompanyValue, setFormCompanyValue] = useState("");
  const [formContactValue, setFormContactValue] = useState("");
  const [formCountryValue, setFormCountryValue] = useState("");

  const [usersValues, setUsersValues] = useState<Users>({
    firstname: "",
    lastname: "",
    username: "",
    email: "",
    description: "",
  });

  const [users, setUsers] = useState<Users[]>();

  useEffect(() => {
    setTimeout(async () => {
      const response = await fetch("http://localhost:3001/api/v1/count/fetch", {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }).then((res) => {
        return res.json();
      });
      if (response && !response.error) {
        setCount(response.data.count);
      }
    });
  }, []);

  return (
    <div className="p-10">
      <div className="border-2 rounded-md p-5 min-w-max text-center mb-5">
        <p className="text-lg">
          Click to change the count number{" "}
          {"(Count number is stored and handle by the server)"}
        </p>
        <button
          className="bg-slate-600 text-white hover:bg-slate-500 transition duration-300 p-4 rounded-md"
          onClick={() => handleCount(count, setCount)}
        >
          Count is {count}
        </button>
      </div>

      <div className="border-2 rounded-md p-5 min-w-max overflow-x-auto relative mb-5">
        <p className="text-lg text-center pb-5">Company table</p>
        <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
          <thead className="text-xs text-gray-700 uppercase bg-gray-200 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th scope="col" className="py-3 px-6">
                Company
              </th>
              <th scope="col" className="py-3 px-6">
                Contact
              </th>
              <th scope="col" className="py-3 px-6">
                Country
              </th>
            </tr>
          </thead>
          <tbody>
            {companys &&
              companys.map((company) => (
                <tr
                  key={company.company}
                  className="bg-white border-b dark:bg-gray-800 dark:border-gray-700"
                >
                  <th
                    scope="row"
                    className="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white"
                  >
                    {company.company}
                  </th>
                  <td className="py-4 px-6">{company.contact}</td>
                  <td className="py-4 px-6">{company.country}</td>
                </tr>
              ))}
          </tbody>
        </table>
        <div className="m-5">
          <div className="mb-6">
            <label
              htmlFor="company"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Company
            </label>
            <input
              type="text"
              id="company"
              onChange={(e) => setFormCompanyValue(e.target.value)}
              value={formCompanyValue}
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="Alfreds Flutterkiste"
              required
            />
          </div>
          <div className="mb-6">
            <label
              htmlFor="contact"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Contact
            </label>
            <input
              type="text"
              id="contact"
              onChange={(e) => setFormContactValue(e.target.value)}
              value={formContactValue}
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="Maria Anders"
              required
            />
          </div>
          <div className="mb-6">
            <label
              htmlFor="country"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Country
            </label>
            <input
              type="text"
              id="country"
              onChange={(e) => setFormCountryValue(e.target.value)}
              value={formCountryValue}
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="Germany"
              required
            />
          </div>
          <button
            onClick={() =>
              handleForm(
                formCompanyValue,
                formContactValue,
                formCountryValue,
                setCompanys
              )
            }
            className="text-white bg-slate-600 hover:bg-slate-800 transition duration-200 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Insert
          </button>
        </div>
      </div>

      <div className="border-2 rounded-md p-5 min-w-max overflow-x-auto relative mb-5">
        <p className="text-lg text-center pb-5">Users</p>
        <div className="m-5">
          <div className="mb-6">
            <label
              htmlFor="firstname"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Firstname
            </label>
            <input
              type="text"
              id="firstname"
              onChange={(e) => {
                setUsersValues({ ...usersValues, firstname: e.target.value });
              }}
              value={usersValues.firstname}
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="John"
              required
            />
          </div>
          <div className="mb-6">
            <label
              htmlFor="lastname"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Lastname
            </label>
            <input
              type="text"
              id="lastname"
              onChange={(e) => {
                setUsersValues({ ...usersValues, lastname: e.target.value });
              }}
              value={usersValues.lastname}
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="Smith"
              required
            />
          </div>
          <div className="mb-6">
            <label
              htmlFor="username"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Username
            </label>
            <input
              type="text"
              id="username"
              onChange={(e) => {
                setUsersValues({ ...usersValues, username: e.target.value });
              }}
              value={usersValues.username}
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="Windows 11"
              required
            />
          </div>
          <div className="mb-6">
            <label
              htmlFor="email"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Email
            </label>
            <input
              type="text"
              id="email"
              onChange={(e) => {
                setUsersValues({ ...usersValues, email: e.target.value });
              }}
              value={usersValues.email}
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="example@example.com"
              required
            />
          </div>
          <div className="mb-6">
            <label
              htmlFor="description"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
            >
              Description
            </label>
            <input
              type="text"
              id="description"
              onChange={(e) => {
                setUsersValues({ ...usersValues, description: e.target.value });
              }}
              value={usersValues.description}
              className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              placeholder="And then there was..."
              required
            />
          </div>
          <button
            onClick={() => {
              createUser(usersValues);
            }}
            className="text-white bg-slate-600 hover:bg-slate-800 transition duration-200 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Add
          </button>

          <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400 mt-5">
            <thead className="text-xs text-gray-700 uppercase bg-gray-200 dark:bg-gray-700 dark:text-gray-400">
              <tr>
                <th scope="col" className="py-3 px-6">
                  Firstname
                </th>
                <th scope="col" className="py-3 px-6">
                  Lastname
                </th>
                <th scope="col" className="py-3 px-6">
                  Username
                </th>
                <th scope="col" className="py-3 px-6">
                  Email
                </th>
                <th scope="col" className="py-3 px-6">
                  Description
                </th>
              </tr>
            </thead>
            <tbody>
              {users &&
                users.map((user) => (
                  <tr
                    key={user.email}
                    className="bg-white border-b dark:bg-gray-800 dark:border-gray-700"
                  >
                    <th
                      scope="row"
                      className="py-4 px-6 font-medium text-gray-900 whitespace-nowrap dark:text-white"
                    >
                      {user.firstname}
                    </th>
                    <td className="py-4 px-6">{user.lastname}</td>
                    <td className="py-4 px-6">{user.username}</td>
                    <td className="py-4 px-6">{user.email}</td>
                    <td className="py-4 px-6">{user.description}</td>
                  </tr>
                ))}
            </tbody>
          </table>
          <button
            onClick={() => {
              refreshUsers(users, setUsers);
            }}
            className="text-white mt-5 bg-slate-600 hover:bg-slate-800 transition duration-200 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Refresh
          </button>
        </div>
      </div>
    </div>
  );
}
