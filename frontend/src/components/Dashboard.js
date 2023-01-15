import Header from "./Header";
import * as React from 'react';
import axios from 'axios';
import {
    useNavigate,
} from "react-router-dom";

export default function Dashboard() {

    const [login_checking, setLoginChecking] = React.useState(true);
    var navigate = useNavigate();

    const [data, setData] = React.useState([]);
    console.log(data);

    async function GetRepos(e) {
        await axios.post(`${process.env.REACT_APP_BACKEND_URL}/api/get-repos`, {
            "installation_id": e
        }).then((res) => {
            setData(res.data);
        }).catch((err) => {
            console.log(err);
        });
    }

    const [apps, setApps] = React.useState([]);
    async function GetAllApps() {
        //Get request to backend to check if user is logged in
        await axios.get(`${process.env.REACT_APP_BACKEND_URL}/api/installed-apps`).then((res) => {
            setApps(res.data);
        }).catch((err) => {
            console.log(err);
        });
    }

    async function CheckLogin() {
        //Get request to backend to check if user is logged in
        await axios.get(`${process.env.REACT_APP_BACKEND_URL}/api/is-logged-in`).then((res) => {
            if (res.data.message !== "Authorized") {
                navigate("/login");
            }
            setLoginChecking(false);
        }).catch((err) => {
            console.log(err);
            setLoginChecking(false);
            navigate("/login");
        });
    }

    React.useEffect(() => {
        GetAllApps();
        CheckLogin();
    }, []);

    const [isloading, setIsLoading] = React.useState(false);
    const [repos, setRepos] = React.useState([]);

    console.log(repos);

    async function getAllRepos() {
        //Set loading to true
        setIsLoading(true);
        //Axios get request with header 
        await axios.get(`${process.env.REACT_APP_BACKEND_URL}/api/get-my-repos`).then((res) => {
            setRepos(res.data.repositories);
            setIsLoading(false);
        }).catch((err) => {
            console.log(err);
            setIsLoading(false);
        });
    }


    React.useEffect(() => {
        getAllRepos();
    }, []);

    return (
        <div>
            <Header />
            <div className="p-5 hero">
                <div className="hero-content min-w-full items-start">
                    <div className="w-[30%] h-[85%] justify-center">
                        <h1 className="text-4xl font-bold">Welcome to GitCode</h1>
                        <p className="text-lg mt-5">
                            Select a repository to view its details and to view its commits.
                        </p>
                        <select className="select select-bordered w-full max-w-xs mt-5" onChange={
                            (e) => {
                                GetRepos(e.target.value);
                            }}>
                            <option disabled defaultValue>Select App And View Repos</option>
                            {apps.map((app, index) => {
                                return (
                                    <option onClick={(e) => {
                                        GetRepos(app.installation_id);
                                    }} key={index} value={app.installation_id}>{app.installation_id}</option>
                                )
                            })}
                        </select>
                    </div>
                    <div className="overflow-x-scroll h-[85vh] w-[70%]">
                        <table className="table w-full">
                            {/* <!-- head --> */}
                            <thead>
                                <tr>
                                    <th></th>
                                    <th>Name</th>
                                    <th>Visibility</th>
                                    <th>Created AT</th>
                                    <th>Details</th>
                                </tr>
                            </thead>
                            <tbody>
                                {isloading ? <div className="flex w-full h-full justify-center items-center"><div>Loading...</div></div> : repos !== null && repos.map((repo, index) => {
                                    return (
                                        <tr key={index}>
                                            <th>
                                                <label>
                                                    {/* Index */}
                                                    {index + 1}

                                                </label>
                                            </th>
                                            <td>
                                                <div className="flex items-center space-x-3">
                                                    <div className="avatar">
                                                        <div className="mask mask-squircle w-12 h-12">
                                                            <img src={repo.owner.avatar_url} alt="Avatar Tailwind CSS Component" />
                                                        </div>
                                                    </div>
                                                    <div>
                                                        <div className="font-bold">{repo.full_name}</div>
                                                        <div className="text-sm opacity-50">{repo.language}</div>
                                                    </div>
                                                </div>
                                            </td>
                                            <td>
                                                {repo.visibility !== "public" ? <span className="badge badge-error">Private</span> : <span className="badge badge-success">Public</span>}
                                                <br />
                                                <span className="badge badge-ghost badge-sm">
                                                    {/* Last update in ago */}
                                                    {repo.default_branch !== null ? repo.default_branch : "No update"}
                                                </span>
                                            </td>
                                            <td>
                                                <div className="text-sm opacity-50">{repo.created_at}</div>
                                            </td>
                                            <th>
                                                <button className="btn btn-ghost btn-xs"
                                                onClick={(e) => {
                                                        // Onclick copy to clipboard
                                                        navigator.clipboard.writeText(repo.clone_url);
                                                    }}>Copy Clone URL</button>
                                            </th>
                                        </tr>
                                    )
                                })}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </div>
    )
}