import { useEffect, useState } from "react";
import { useRatio } from "./context/RatioContext";
import { useUser } from "./context/UserContext";
import { Button, HStack, Input, PinInput, PinInputField, Spinner } from "@chakra-ui/react";

const Buy = () => {
    const { user } = useUser();
    const { bearer } = useRatio();
    const { ratio } = useRatio();
    const { phoneId } = useRatio();
    const { triedLogin } = useRatio();
    const { freezeUser } = useRatio();
    const { sendPhone } = useRatio();
    const { reSendPhone } = useRatio();
    const { sendOtp } = useRatio();
    const { initializeRatio } = useRatio();
    const { resetTriedLogin } = useRatio();
    const { resetRatioState } = useRatio();
    const [userPhone, setUserPhone] = useState('');
    const [otp, setOtp] = useState('');

    const { sendPhonePending } = useRatio();
    const { sendPhoneError } = useRatio();
    const { sendPhoneErrorMessage } = useRatio();

    const { sendOtpPending } = useRatio();
    const { sendOtpError } = useRatio();
    const { sendOtpErrorMessage } = useRatio();

    const { initializeRatioPending } = useRatio();
    const { initializeRatioError } = useRatio();
    const { initializeRatioErrorMessage } = useRatio();

    
    

    useEffect(() => {
        resetTriedLogin();
        if(user != freezeUser){
            resetRatioState();
            initializeRatio();
        }else if(!bearer){
            initializeRatio();
        }
    }, [])

    



    return (
        <div className="buy">
            {!bearer && triedLogin && !initializeRatioPending && <Button 
                className="button" 
                onClick={()=>{
                    initializeRatio();
                    }}
                >Log In with Ratio
            </Button>}
            {initializeRatioPending && <Spinner size="lg"/>}
            {bearer && !phoneId && !sendPhonePending && <div>
                <form onSubmit={ async (e) => {
                    e.preventDefault();
                    const phone = '+' + userPhone;
                    console.log(phone);
                    sendPhone(phone);
                    console.log(ratio);
                }}>
                <h1>Enter Phone Number</h1>
                <Input 
                    type="tel" 
                    placeholder="Please Include Country Code" 
                    onChange={(e) => setUserPhone(e.target.value)} 
                    value={userPhone}/>
                <br/>
                <Button 
                    type='submit' 
                    className="button"
                    >Submit
                </Button>
                </form>
            </div>}
            {sendPhonePending && <Spinner size="lg"/>}
            {!ratio && phoneId &&  !sendOtpPending && <div>
                <form onSubmit={(e) => {
                    e.preventDefault();
                    sendOtp(otp);
                }}>
                    <h1>Enter your One Time Passcode</h1>
                    <div className="pin">
                        <PinInput 
                            otp 
                            type="number"   
                            value={otp} 
                            onChange={
                                (ev) => {
                                    setOtp(ev);
                                    if(ev.length==6){
                                        sendOtp(ev);
                                        setOtp('');            
                                    }
                                }
                            }>
                            <PinInputField className="pinInput"/>
                            <PinInputField className="pinInput"/>
                            <PinInputField className="pinInput"/>
                            <PinInputField className="pinInput"/>
                            <PinInputField className="pinInput"/>
                            <PinInputField className="pinInput"/>
                        </PinInput>
                    </div>
                    <br/>
                    <Button 
                        type='submit' 
                        className="button"
                        >Submit
                    </Button>
                </form>
                <Button onClick={() => {reSendPhone()}}>Resend OTP</Button>
            </div>}
            {sendOtpPending && <Spinner size="lg"/>}
            {ratio && <h1>Logged in with Ratio</h1>}
        </div>);
}
 
export default Buy;