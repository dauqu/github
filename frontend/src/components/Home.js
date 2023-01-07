import Header from "./Header";
import * as React from 'react';
import axios from 'axios';

export default function Home() {



    const [isloading, setIsLoading] = React.useState(false);
    const [repos, setRepos] = React.useState([]);

    async function getAllRepos() {
        //Set loading to true
        setIsLoading(true);
        //Axios get request with header 
        await axios.get("http://localhost:8000/api/get-my-repos").then((res) => {
            setRepos(res.data.repositories);
            setIsLoading(false);
        }).catch((err) => {
            console.log(err);
            setIsLoading(false);
        });
    }

    console.log(repos);

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
                        <p className="text-lg">GitCode is a web application that allows you to search for a GitHub user and view their repositories.</p>

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
                                </tr>
                            </thead>
                            <tbody>
                                {isloading ? <div className="flex w-full h-full justify-center items-center"><div>Loading...</div></div> : repos !== null && repos.map((repo, index) => {
                                    return (
                                        <tr key={index}>
                                            <th>{index + 1}</th>
                                            <td>{repo.full_name}</td>
                                            <td>
                                                {repo.visibility !== "public" ? <span className="badge badge-error">Private</span> : <span className="badge badge-success">Public</span>}
                                            </td>
                                            <td>{repo.created_at}</td>
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