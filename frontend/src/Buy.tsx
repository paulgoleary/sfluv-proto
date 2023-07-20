import { useEffect, useState } from "react";
import { useRatio } from "./context/RatioContext";
import { useUser } from "./context/UserContext";
import { Button, Checkbox, HStack, Input, PinInput, PinInputField, Spinner } from "@chakra-ui/react";
import 'react-phone-number-input/style.css'
import PhoneInput, { isValidPhoneNumber } from 'react-phone-number-input'
import { parsePhoneNumber } from "react-phone-number-input";


const Buy = () => {
    const { user } = useUser();
    const { bearerValue } = useRatio();
    const { ratio } = useRatio();
    const { phoneId } = useRatio();
    const { userId } = useRatio();
    const { triedLogin } = useRatio();
    const { freezeUser } = useRatio();
    const { userSubmitted } = useRatio();
    const { userPhoneContext } = useRatio();
    const { kyc } = useRatio();
    const { sendPhone } = useRatio();
    const { reSendPhone } = useRatio();
    const { sendOtp } = useRatio();
    const { sendUser } = useRatio();
    const { sendKyc } = useRatio();
    const { setUserData } = useRatio();
    const { setKycData } = useRatio();
    const { initializeRatio } = useRatio();
    const { resetTriedLogin } = useRatio();
    const { resetRatioState } = useRatio();
    const [ userPhone, setUserPhone ] = useState<string | null>();
    const [ firstName, setFirstName ] = useState('');
    const [ middleName, setMiddleName ] = useState('');
    const [ lastName, setLastName ] = useState('');
    const [ email, setEmail ] = useState('');
    const [ acceptedTerms, setAcceptedTerms ] = useState(false);
    const [ phoneValid, setPhoneValid ] = useState(true);
    const [ otp, setOtp ] = useState('');
    const [ testPhone, setTestPhone ] = useState('');

    const [ dateOfBirth, setDateOfBirth ] = useState('');
    const [ idType, setIdType ] = useState('');
    const [ idNumber, setIdNumber ] = useState('');
    const [ line1, setLine1 ] = useState('');
    const [ city, setCity ] = useState('');
    const [ state, setState ] = useState('');
    const [ postalCode, setPostalCode ] = useState('');


    const { sendPhonePending } = useRatio();
    const { sendPhoneError } = useRatio();
    const { sendPhoneErrorMessage } = useRatio();

    const { sendOtpPending } = useRatio();
    const { sendOtpError } = useRatio();
    const { sendOtpErrorMessage } = useRatio();

    const { initializeRatioPending } = useRatio();
    const { initializeRatioError } = useRatio();
    const { initializeRatioErrorMessage } = useRatio();

    const { sendKycPending } = useRatio();

    
    

    useEffect(() => {
        resetTriedLogin();
        if(user != freezeUser){
            resetRatioState();
            initializeRatio();
        }else if(!bearerValue){
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
        <div className="buy content">
            {!bearerValue && triedLogin && !initializeRatioPending && <Button 
                className="button" 
                onClick={()=>{
                    initializeRatio();
                    }}
                >Log In with Ratio
            </Button>}
            {initializeRatioPending && <Spinner size="lg"/>}
            {bearerValue && !phoneId && !sendPhonePending && <div>
                <form className="form"onSubmit={ async (e) => {
                    e.preventDefault();
                    const phone = '' + userPhone;
                    if(isValidPhoneNumber(phone)){
                        setPhoneValid(true);
                        const country  = ('' + parsePhoneNumber(phone)?.country).toString();
                        console.log(country);
                        console.log(phone);
                        sendPhone(phone, bearerValue);
                        setUserData(
                            firstName, 
                            middleName, 
                            lastName, 
                            email, 
                            country, 
                            phone, 
                            acceptedTerms);
                        console.log(ratio);
                    }else{
                        setPhoneValid(false);
                    }
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
                <PhoneInput
                    required
                    placeholder="Enter phone number"
                    value={userPhone ? userPhone : ''}
                    onChange={setUserPhone}/>
                    {!phoneValid && <p className="error">Phone Number Invalid</p>}
                <Checkbox 
                    required
                    onChange={(e) => setAcceptedTerms(e.target.checked)} 
                    >I agree to the&nbsp;
                    <a href="https://www.ratio.me/legal/ratio-labs-usa-inc-terms-of-service" target="_blank">Terms of Service</a>
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
                    <h2>Sent to the Phone Number Ending in {!userPhone && <span>{userPhoneContext?.substring(8)}</span>}{userPhone?.substring(8)}</h2>
                    <br/>
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
            {sendKycPending && <Spinner size="lg"/>}
            {ratio && userId && !kyc && !sendKycPending && <div>
                <form onSubmit={(e) => {
                    e.preventDefault();
                    const idType = 'SSN';
                    console.log(dateOfBirth)
                    setKycData( dateOfBirth, idType, idNumber, line1, city, state, postalCode );
                    sendKyc( dateOfBirth, idType, idNumber, line1, city, state, postalCode );
                }}>
                    <h1>Verify your Identity</h1>
                    <label>Date of Birth:</label>
                    <Input 
                        type="date" 
                        required
                        placeholder="Date of Birth" 
                        onChange={(e) => setDateOfBirth(e.target.value)} 
                        value={dateOfBirth}/>
                    <Input 
                        type="text" 
                        required
                        placeholder="SSN"
                        minLength={9}
                        maxLength={9}
                        onChange={(e) => setIdNumber(e.target.value)} 
                        value={idNumber}/> 
                    <Input 
                        type="text" 
                        required
                        placeholder="Street Address" 
                        onChange={(e) => setLine1(e.target.value)} 
                        value={line1}/> 
                    <Input 
                        type="text" 
                        required
                        placeholder="City" 
                        onChange={(e) => setCity(e.target.value)} 
                        value={city}/>  
                    <Input 
                        type="text" 
                        required
                        placeholder="State" 
                        onChange={(e) => setState(e.target.value)} 
                        value={state}/>
                    <Input 
                        type="text" 
                        required
                        placeholder="Postal Code" 
                        minLength={5}
                        onChange={(e) => setPostalCode(e.target.value)} 
                        value={postalCode}/>
                    <Button 
                        type='submit' 
                        className="button"
                        >Submit
                    </Button>
                </form>
            </div>}
            {(kyc == 'IN_REVIEW') && <div>
                <h1>Verification Submitted</h1>
                <h2>You will be recieving an email with your verification status shortly.</h2>
                </div>}
            {(kyc == 'APPROVED') && <div>
                <h1>Verification Approved</h1>
                <h2>You may now purchase SFLUV.</h2>
                </div>}
        </div>);
}
 
export default Buy;