import { useEffect, useState } from "react";
import { useRatio } from "./context/RatioContext";
import { useUser } from "./context/UserContext";
import { Button, Checkbox, HStack, Input, PinInput, PinInputField, Spinner } from "@chakra-ui/react";

const Buy = () => {
    const { user } = useUser();
    const { bearer } = useRatio();
    const { ratio } = useRatio();
    const { phoneId } = useRatio();
    const { triedLogin } = useRatio();
    const { freezeUser } = useRatio();
    const { userSubmitted } = useRatio();
    const { sendPhone } = useRatio();
    const { reSendPhone } = useRatio();
    const { sendOtp } = useRatio();
    const { sendUser } = useRatio();
    const { setUserData } = useRatio();
    const { initializeRatio } = useRatio();
    const { resetTriedLogin } = useRatio();
    const { resetRatioState } = useRatio();
    const [ userPhone, setUserPhone ] = useState('');
    const [ firstName, setFirstName ] = useState('');
    const [ middleName, setMiddleName ] = useState('');
    const [ lastName, setLastName ] = useState('');
    const [ email, setEmail ] = useState('');
    const [ country, setCountry ] = useState('');
    const [ acceptedTerms, setAcceptedTerms ] = useState(false);
    const [ otp, setOtp ] = useState('');
    const [ testPhone, setTestPhone ] = useState('');

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
    }, []);

    useEffect(() => {
        if(ratio && userSubmitted){
            console.log("Sent User Data!")
            sendUser();
        }
    }, [ratio]);


    



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
                    setUserData(
                        firstName, 
                        middleName, 
                        lastName, 
                        email, 
                        country, 
                        phone, 
                        acceptedTerms);
                    console.log(ratio);
                }}>
                <h1>Sign Up with Ratio</h1>
                <Input 
                    type="text" 
                    required
                    placeholder="First Name" 
                    onChange={(e) => setFirstName(e.target.value)} 
                    value={firstName}/>
                <Input 
                    type="text" 
                    required
                    placeholder="Middle Name" 
                    onChange={(e) => setMiddleName(e.target.value)} 
                    value={middleName}/>
                <Input 
                    type="text" 
                    required
                    placeholder="Last Name" 
                    onChange={(e) => setLastName(e.target.value)} 
                    value={lastName}/>
                <Input 
                    type="text" 
                    required
                    placeholder="Email" 
                    onChange={(e) => setEmail(e.target.value)} 
                    value={email}/>
                <Input 
                    type="text" 
                    required
                    placeholder="Country" 
                    onChange={(e) => setCountry(e.target.value)} 
                    value={country}/>
                <Input 
                    type="tel" 
                    required
                    placeholder="Please Include Country Code" 
                    onChange={(e) => setUserPhone(e.target.value)} 
                    value={userPhone}/>
                <Checkbox 
                    required
                    onChange={(e) => setAcceptedTerms(e.target.checked)} 
                    >I agree to the Terms and Conditions
                </Checkbox>
                <br/><br/>
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