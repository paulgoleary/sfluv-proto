import { useEffect, useState } from "react";
import { useUser } from "./context/UserContext";
import links from "./links";


const Test = () => {
    const { luv_server } = links();
    const [data, setData] = useState([]);
    const walletAddress = useUser().user;
    const walletNetwork = "POLYGON";
    const [signedData, setSignedData] = useState('');
    const inputs = { walletAddress, walletNetwork };
    const inputs2 = { walletAddress, walletNetwork, signedData };
    const inputData = {
        method: 'POST',
        headers: {'Content-Type': 'application/json', 'Accept': 'application/json'},
        body: JSON.stringify(inputs)
    };
    const inputData2 = {
        method: 'POST',
        headers: {'Content-Type': 'application/json', 'Accept': 'application/json'},
        body: JSON.stringify(inputs2)
    };

    const fetchData = async ( link:string, postData:object ) => {
        console.log(inputs);
        const response = await fetch(link, postData);
        console.log(inputs);
        if (!response.ok) {
            throw new Error('Data could not be fetched!')
          } else {
            return response.json()
          }

          
    }
    useEffect(() => {
        fetchData( luv_server + '/ratio/sessions' , inputData )
        .then((res) => {;
            setData(res);
        })
        .catch((e) => {
            console.log(e.message);
        })
    }, []);
    return (
        <div>
            <h1>Test</h1>
            <span>{data}</span>
        </div>
    );
}
 
export default Test;