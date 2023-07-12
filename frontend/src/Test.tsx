import { useEffect, useState } from "react";
import { useUser } from "./context/UserContext";
import links from "./links";
import { magic } from "./libs/magic";


const Test = () => {
    const [data, setData] = useState('');
    
    return (
        <div>
            <h1>Test</h1>
            <span>{data}</span>
        </div>
    );
}
 
export default Test;