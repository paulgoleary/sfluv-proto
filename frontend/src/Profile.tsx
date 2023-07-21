import { Route, Router, Routes } from "react-router-dom";
import MagicWallet from "./MagicWallet";
import RatioProfile from "./RatioProfile";
import { useState } from "react";
import { useRatio } from "./context/RatioContext";

const Profile = () => {
    
    const [ ratioSelected, setRatioSelected ] = useState(false);
    const [ magicWalletSelected, setMagicWalletSelected ] = useState(true);

    const { firstName } = useRatio();


    return (
    <div className="profile">
        { firstName && 
            <div className="profileNav">
                {magicWalletSelected && <button 
                    onClick={(e) => {
                        setRatioSelected(false);
                        setMagicWalletSelected(true);
                    }}
                    className="magic selected"
                    >Magic Wallet
                </button>}
                {ratioSelected && <button 
                    onClick={(e) => {
                        setRatioSelected(false);
                        setMagicWalletSelected(true);
                    }}
                    className="magic"
                    >Magic Wallet
                </button>}
                {ratioSelected && <button 
                    onClick={(e) => {
                        setMagicWalletSelected(false);
                        setRatioSelected(true);
                    }}
                    className="ratio selected"
                    >Ratio Profile
                </button>}
                {magicWalletSelected && <button 
                    onClick={(e) => {
                        setMagicWalletSelected(false);
                        setRatioSelected(true);
                    }}
                    className="ratio"
                    >Ratio Profile
                </button>}
            </div>
        }
        <div>
            { ratioSelected && <RatioProfile/> }
            { magicWalletSelected && <MagicWallet/> }
        </div>
     </div>);
}
 
export default Profile;