import axios from 'axios';

export default function GitCode() {
    //get data from query string
    const query = new URLSearchParams(window.location.search);
    const code = query.get("installation_id");

    console.log(code);

    //Post request to backend
    async function postData() {
        await axios.post(`${process.env.REACT_APP_BACKEND_URL}/api/install-app`, {
            installation_id: code
        }).then((response) => {
            console.log(response);
            // window.close();
        }).catch((error) => {
            console.log(error);
            // window.close();
        });
    }

    postData();

    return (
        <div>
        </div>
    )
}