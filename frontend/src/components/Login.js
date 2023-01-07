import * as React from "react";
import axios from "axios";


export default function Login() {

    const [email, setEmail] = React.useState("");
    const [password, setPassword] = React.useState("");
    const [isloading, setIsLoading] = React.useState(false);

    async function login() {
        //Set loading to true
        setIsLoading(true);
        //Axios Post request to backend
        await axios.post("http://localhost:8000/api/login", {
            Email: email,
            Password: password,
        }).then((res) => {
            console.log(res);
            //Set loading to false
            setIsLoading(false);
            //If login is successful, redirect to home page
            window.location.href = "/";
        }).catch((err) => {
            console.log(err);
            //Set loading to false
            setIsLoading(false);
            alert("Login Failed");
        });
    }

    return (
        <div>
            <div className="hero min-h-screen bg-base-200 rounded-none">
                <div className="hero-content flex-col">
                    <div className="text-center">

                    </div>
                    <div className="card flex-shrink-0 w-full shadow-sm bg-base-100 rounded-none min-w-[500px]">
                        <div className="card-body">
                            <div className="form-control">
                                <label className="label">
                                    <span className="label-text">Email</span>
                                </label>
                                <input type="text" placeholder="email" className="input input-bordered"
                                    value={email}
                                    onChange={(e) => {
                                        setEmail(e.target.value);
                                    }} />
                            </div>
                            <div className="form-control">
                                <label className="label">
                                    <span className="label-text">Password</span>
                                </label>
                                <input type="text" placeholder="password" className="input input-bordered"
                                    value={password}
                                    onChange={(e) => {
                                        setPassword(e.target.value);
                                    }} />
                                <label className="label">
                                    <a href="#" className="label-text-alt link link-hover">Forgot password?</a>
                                </label>
                            </div>
                            <div className="form-control mt-6">
                                <button className={`btn btn-primary ${isloading ? "loading" : ""}`} onClick={() => {
                                    login();
                                }}>Login</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}