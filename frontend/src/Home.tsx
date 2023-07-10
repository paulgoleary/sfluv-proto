import { useState } from "react";
import { useUser } from "./context/UserContext";
import useBalance from "./useBalance";
import links from "./links";

const Home = () => {
    const { luv_server } = links();
    const { user } = useUser();
    const balance = useBalance();

    return ( 
        <div className="home">
            <h1>{user}</h1>
            {user && <h2>Balance: {balance}</h2>}
        </div> 
    );
}
 
export default Home;