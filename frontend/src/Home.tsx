import { useState } from "react";
import { useUser } from "./context/UserContext";
import useBalance from "./useBalance";

const Home = () => {
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