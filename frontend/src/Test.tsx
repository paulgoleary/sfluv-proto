import { useEffect, useState } from "react";
import { useUser } from "./context/UserContext";
import links from "./links";
import { magic } from "./libs/magic";
import { useWeb3 } from "./context/Web3Context";



const Test = () => {
    const [data, setData] = useState('');
    const { web3 } = useWeb3();
    const { user } = useUser();
    return (
        <div className="test content">
            <h1>Test</h1>
            <span>{data}</span>
        </div>
    );
}
 
export default Test;