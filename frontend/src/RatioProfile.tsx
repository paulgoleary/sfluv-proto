import { useEffect, useState } from "react";
import { useRatio } from "./context/RatioContext";
import { PhoneNumber } from "libphonenumber-js/types";

const RatioProfile = () => {
    const { firstName } = useRatio();
    const { middleName } = useRatio();
    const { lastName } = useRatio();
    const { email } = useRatio();
    const { phone } = useRatio();
    const { kyc } = useRatio();

    const [ tempFirst, setTempFirst ] = useState('');
    const [ tempMiddle, setTempMiddle ] = useState('');
    const [ tempLast, setTempLast ] = useState('');
    const [ tempEmail, setTempEmail ] = useState('');
    const [ tempPhone, setTempPhone ] = useState('');

    const [ editModeOn, setEditModeOn ] = useState(false);

    useEffect(() => {
        if(firstName){
            setTempFirst(firstName);
        }
        if(middleName){
            setTempMiddle(middleName);
        }
        if(lastName){
            setTempLast(lastName);
        }
        if(email){
            setTempEmail(email);
        }
        if(phone){
            setTempPhone(phone);
        }
    }, [ firstName, middleName, lastName, email, phone ])

    return (
        <div>
            <form className="profileSection">
                <h2 className="profileSectionTitle">Basic Info</h2>
                <div className="profileItem firstName">
                    <div className="profileItemCategory"><a>First Name:</a></div>
                    <div className="profileItemValue">{firstName && <input 
                        type="text" 
                        value={tempFirst}
                        onChange={(e) => {
                            setTempFirst(e.target.value);
                        }}
                    />}</div>
                </div>
                <div className="profileItem middleName">
                    <div className="profileItemCategory"><a>Middle Name:</a></div>
                    <div className="profileItemValue">{middleName && <input 
                        type="text" 
                        value={tempMiddle}
                        onChange={(e) => {
                            setTempMiddle(e.target.value);
                        }}
                    />}</div>
                </div>
                <div className="profileItem lastName">
                    <div className="profileItemCategory"><a>Last Name:</a></div>
                    <div className="profileItemValue">{lastName && <input 
                        type="text" 
                        value={tempLast}
                        onChange={(e) => {
                            setTempLast(e.target.value);
                        }}
                    />}</div>
                </div>
                <div className="profileItem phone">
                    <div className="profileItemCategory"><a>Phone Number:</a></div>
                    <div className="profileItemValue">{phone && <input 
                        type="text" 
                        value={tempPhone}
                        onChange={(e) => {
                            setTempPhone(e.target.value);
                        }}
                    />}</div>
                </div>
                <div className="lastProfileItem email">
                    <div className="profileItemCategory"><a>Email:</a></div>
                    <div className="profileItemValue">{email && <input 
                        type="text" 
                        value={tempEmail}
                        onChange={(e) => {
                            setTempEmail(e.target.value);
                        }}
                    />}</div>
                </div>
            </form>
            <div className="profileSection">
                <h2 className="profileSectionTitle">Verification Info</h2>
                <div className="lastProfileItem">
                    <div className="profileItemCategory">Status:</div>
                    <div className="profileItemValue">{kyc}{!kyc && <a>Not Started</a>}</div>
                </div>
            </div>
        </div>
    );
}
 
export default RatioProfile;