import { useState } from "react";
import { useUser } from "./context/UserContext";
import useBalance from "./useBalance";
import links from "./links";
import { useRatio } from "./context/RatioContext"
import { Button } from "@chakra-ui/react";

const Home = () => {
    const { luv_server } = links();
    const { user } = useUser();
    const { ratio } = useRatio();
    const { userId } = useRatio();
    const balance = useBalance();

    
    return ( 
        <div className="home content">
            <h1>{user}</h1>
            {user && <h2>Balance: {balance}</h2>}
            {ratio && <h2>Ratio Logon: <div className="ratioLogon"><div className="ratioLogonText">{ratio}</div></div></h2>}
            {userId && <h2>Ratio User ID: {userId}</h2>}
        </div> 
    );
}
 
export default Home;