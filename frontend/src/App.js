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

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
    children: [

    ],
  },
  {
    path: "login",
    element: <Login />,
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
    axios.get("http://localhost:8000/api/is-logged-in").then((res) => {
      setIsLoading(false);
    }).catch((err) => {
      console.log(err);
      setIsLoading(false);
      navigate("/login");
    });
  }

  React.useEffect(() => {
    CheckLogin();
  }, []);

  return (
    <div className="App">
      {isloading ? <div>Loading...</div> : <RouterProvider router={router} />}
    </div>
  );
}

export default App;
