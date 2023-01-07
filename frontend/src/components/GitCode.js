import axios from 'axios';

export default function GitCode() {
    //get data from query string
    const query = new URLSearchParams(window.location.search);
    const code = query.get("installation_id");

    console.log(code);

    //Post request to backend
    async function postData() {
        await axios.post("http://localhost:8000/api/connect-github", {
            installation_id: code
        }).then((response) => {
            console.log(response);
            window.close();
        }).catch((error) => {
            console.log(error);
            window.close();
        });
    }

    postData();

    return (
        <div>
        </div>
    )
}