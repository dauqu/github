import * as React from 'react';
import axios from 'axios';

export default function Header() {

    //get profile data
    const [profile, setProfile] = React.useState({});
    const [isloading, setIsLoading] = React.useState(true);

    async function getProfile() {
        await axios.get("http://localhost:8000/api/profile").then((res) => {
            setProfile(res.data);
            setIsLoading(false);
        }).catch((err) => {
            console.log(err);
            setIsLoading(false);
        });
    }

    // React.useEffect(() => {
    //     getProfile();
    // }, []);


    // const url = "https://github.com/login/oauth/authorize?client_id=Iv1.72f299b0ba45be0a&redirect_uri=http://localhost:3000/gitcode/&scope=user,repo";
        const url = "https://github.com/apps/dauqu/installations/new";

    const openWindow = () => {
        const width = 600;
        const height = 600;
        const left = window.screen.width / 2 - width / 2;
        const top = window.screen.height / 2 - height / 2;
        const windowFeatures = `toolbar=no, menubar=no, width=${width}, height=${height}, top=${top}, left=${left}`;
        const newWindow = window.open(url, "Github Login", windowFeatures);
        if (window.focus) {
            newWindow.focus();
        }
    }

    return (
        <div>
            <div className="navbar shadow-md bg-slate-800">
                <div className="flex-1">
                    <a className="btn btn-ghost normal-case text-xl text-white">Dauqu</a>
                </div>
                <div className="flex-2">
                    <button className="btn btn-sm mr-5" onClick={
                        () => {
                            //Open subwindow
                            openWindow();
                        }
                    }>Connect to Github</button>
                </div>
                <div className="flex-none">
                    <div className="dropdown dropdown-end">
                        <label tabIndex={0} className="btn btn-ghost btn-circle">
                            <div className="indicator fill-white">
                                <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="white" viewBox="0 0 24 24" stroke="currentColor"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" /></svg>
                                <span className="badge badge-sm indicator-item">8</span>
                            </div>
                        </label>
                        <div tabIndex={0} className="mt-3 card card-compact dropdown-content w-52 bg-base-100 shadow">
                            <div className="card-body">
                                <span className="font-bold text-lg">8 Items</span>
                                <span className="text-info">Subtotal: $999</span>
                                <div className="card-actions">
                                    <button className="btn btn-primary btn-block">View cart</button>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="dropdown dropdown-end">
                        <label tabIndex={0} className="btn btn-ghost btn-circle avatar">
                            <div className="w-10 rounded-full">
                                <img src="https://placeimg.com/80/80/people" />
                            </div>
                        </label>
                        <ul tabIndex={0} className="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52">
                            <li>
                                <a className="justify-between">
                                    Profile
                                    <span className="badge">New</span>
                                </a>
                            </li>
                            <li><a>Settings</a></li>
                            <li><a>Logout</a></li>
                        </ul>
                    </div>
                </div>
            </div>
        </div>
    )
}