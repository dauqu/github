import './App.css';
import * as React from "react";
import Login from './components/Login';
import axios from "axios";

import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import Home from './components/Home';
import GitCode from './components/GitCode';
import Register from './components/Register';

const router = createBrowserRouter([
  {
    path: "/home",
    element: <Home />,
    children: [

    ],
  },
  {
    path: "/",
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

  const [isloading, setIsLoading] = React.useState(true);

  async function CheckLogin() {

    //UseNavigate 
    const navigate = router.navigate;

    //Get request to backend to check if user is logged in
    axios.get(`${process.env.REACT_APP_BACKEND_URL}/api/is-logged-in`).then((res) => {
      setIsLoading(false);
      if (res.data.message === "Authorized") {
        navigate("/home");
      }
    }).catch((err) => {
      setIsLoading(false);
    });
  }

  React.useEffect(() => {
    CheckLogin();
  }, []);

  return (
    <div className="">
      {isloading ? <div>Loading...</div> : <RouterProvider router={router} />}
    </div>
  );
}

export default App;
