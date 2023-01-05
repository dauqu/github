
export default function Home() {
    const url = "https://github.com/login/oauth/authorize?client_id=Iv1.72f299b0ba45be0a&redirect_uri=http://localhost:3000/gitcode/&scope=user,repo";

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
            <div className="hero min-h-screen bg-base-200 rounded-none">
                <div className="hero-content flex-col">
                    <div className="text-center">
                        <h1 className="mb-5 text-5xl font-bold">Welcome to React</h1>
                        <p className="max-w-md mb-10 text-2xl">To get started, edit <code>src/App.js</code> and save to reload.</p>
                        <button className="btn btn-primary" onClick={
                            () => {
                                //Open subwindow
                                openWindow();
                            }
                        }>Learn React</button>
                    </div>
                </div>
            </div>
        </div>
    )
}