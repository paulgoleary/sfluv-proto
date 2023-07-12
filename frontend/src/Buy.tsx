import { useEffect, useState } from "react";
import { useRatio } from "./context/RatioContext";
import { useUser } from "./context/UserContext";
import { Button } from "@chakra-ui/react";

const Buy = () => {
    const { user } = useUser();
    const { bearer } = useRatio();
    const { ratio } = useRatio();
    const { phoneId } = useRatio();
    const { sendPhone } = useRatio();
    const { reSendPhone } = useRatio();
    const { sendOtp } = useRatio();
    const { initializeRatio } = useRatio();
    const [userPhone, setUserPhone] = useState('');
    const [phoneSubmitted, setPhoneSubmitted] = useState(false);
    const [otp, setOtp] = useState('');
    const [otpSubmitted, setOtpSubmitted] = useState(false);
    const [phoneSent, setPhoneSent] = useState(false);
    const [initialized, setInitialized] = useState(false);

    useEffect(() => {
        if(!bearer){
            initializeRatio();
        }
    }, [])

    



    return (
        <div className="buy">
            {bearer && !phoneId && <div>
                <form onSubmit={ async (e) => {
                    e.preventDefault();
                    const phone = '+' + userPhone;
                    setPhoneSubmitted(true);
                    console.log(phone);
                    sendPhone(phone);
                    setPhoneSent(true);
                    console.log(ratio);
                }}>
                <h1>Enter Phone Number</h1>
                <input type="tel" placeholder="Please Include Country Code" onChange={(e) => setUserPhone(e.target.value)} value={userPhone}/>
                <br/>
                <Button type='submit' className="button">Submit</Button>
                </form>
            </div>}
            {!ratio && phoneId && <div>
                <form onSubmit={(e) => {
                    e.preventDefault();
                    setOtpSubmitted(true);
                    sendOtp(otp);
                }}>
                    <h1>Enter your One Time Passcode</h1>
                    <input className="otp" type="text" onChange={(e) => setOtp(e.target.value)} value={otp}></input>
                    <br/>
                    <Button type='submit' className="button">Submit</Button>
                </form>
                <Button onClick={() => {reSendPhone()}}>Resend OTP</Button>
            </div>}
            {ratio && <h1>Logged in with Ratio</h1>}
        </div>);
}
 
export default Buy;