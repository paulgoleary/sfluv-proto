import { useEffect, useState } from "react";
import { useRatio } from "./context/RatioContext";
import { useUser } from "./context/UserContext";
import { Button } from "@chakra-ui/react";

const Buy = () => {
    const { user } = useUser();
    const { ratio } = useRatio();
    const { initializeRatio } = useRatio();
    const [userPhone, setUserPhone] = useState('');
    const [phoneSubmitted, setPhoneSubmitted] = useState(false);

    

    



    return (<div>{!phoneSubmitted && <div>
        {/* <form onSubmit={handleSubmit}> */}
        <form onSubmit={(e) => {
        e.preventDefault();
        setUserPhone('+' + userPhone)
        setPhoneSubmitted(true);
        console.log(userPhone);
        initializeRatio(userPhone);
        console.log(ratio);
    }}>
          <h2>Enter Phone Number</h2>
          <input type="tel" placeholder="Please Include Country Code" onChange={(e) => setUserPhone(e.target.value)} value={userPhone}/>
          <button>Submit</button>
        </form>
      </div>}
      {ratio && <h2>{ratio}</h2>}</div>);
}
 
export default Buy;