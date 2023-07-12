import { useState } from "react";
import { useUser } from "./context/UserContext";
import useBalance from "./useBalance";
import links from "./links";
import { useRatio } from "./context/RatioContext"

const Home = () => {
    const { luv_server } = links();
    const { user } = useUser();
    const { ratio } = useRatio();
    const balance = useBalance();

    return ( 
        <div className="home">
            <h1>{user}</h1>
            {user && <h2>Balance: {balance}</h2>}
            {ratio && <h2>Ratio Logon: {ratio}</h2>}
        </div> 
    );
}
 
export default Home;