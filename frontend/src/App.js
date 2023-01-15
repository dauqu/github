import './App.css';
import * as React from "react";
import Login from './components/Login';
import axios from "axios";

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import GitCode from './components/GitCode';
import Register from './components/Register';
import Home from './components/Home';
import Dashboard from './components/Dashboard';
import Profile from './components/Profile';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/dashboard",
    element: <Dashboard />,
    children: [
      {
        path: "profile",
        element: <Profile />,
      },
    ],
  },
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/register",
    element: <Register />,
  },
  {
    path: "/gitcode",
    element: <GitCode />,
  }
]);

//Axios creadentials for all requests
axios.defaults.withCredentials = true;

function App() {

  return (
    <div className="">
      <RouterProvider router={router} />
    </div>
  );
}

export default App;
