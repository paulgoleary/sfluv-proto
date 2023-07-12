import { useEffect, useState } from "react";


const GetFetchData = ( serverLink:string, inputData:object ) => {
    const [data, setData] = useState();
    

    const fetchData = async ( link:string, postData:object ) => {
        console.log('Link: ' + link, 'Post Data: ' + postData);
        const response = await fetch(link, postData);
        if (!response.ok) {
            throw new Error('Data could not be fetched!')
          } else {
            console.log('Response: ' + response.json)
            return response.json()
          }

          
    }
    useEffect(() => {
        fetchData( serverLink , inputData )
        .then((res) => {
            setData(res);
            console.log(data);
        })
        .catch((e) => {
            console.log(e.message);
        })

    }, []);
    return String (data);
}
 
export default GetFetchData;