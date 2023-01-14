import * as React from "react";
import axios from "axios";


export default function Register() {

    const [full_name, setFullName] = React.useState("");
    const [email, setEmail] = React.useState("");
    const [username, setUsername] = React.useState("");
    const [phone, setPhone] = React.useState("");
    const [password, setPassword] = React.useState("");
    const [isloading, setIsLoading] = React.useState(false);

    async function register() {
        //Set loading to true
        setIsLoading(true);
        //Axios Post request to backend
        await axios.post(`${process.env.REACT_APP_BACKEND_URL}/api/register`, {
            FullName: full_name,
            Email: email,
            Username: username,
            Phone: phone,
            LicenseKey: "",
            Password: password
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
                                    <span className="label-text">Full Name</span>
                                </label>
                                <input type="text" placeholder="Full Name" className="input input-bordered"
                                    value={full_name}
                                    onChange={(e) => {
                                        setFullName(e.target.value);
                                    }} />
                            </div>
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
                                    <span className="label-text">Username</span>
                                </label>
                                <input type="text" placeholder="email" className="input input-bordered"
                                    value={username}
                                    onChange={(e) => {
                                        setUsername(e.target.value);
                                    }} />
                            </div>
                            <div className="form-control">
                                <label className="label">
                                    <span className="label-text">Phone</span>
                                </label>
                                <input type="text" placeholder="email" className="input input-bordered"
                                    value={phone}
                                    onChange={(e) => {
                                        setPhone(e.target.value);
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
                                    register();
                                }}>Register</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}