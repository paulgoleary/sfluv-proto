import { Button } from "@chakra-ui/react";
import { useWeb3 } from "./context/Web3Context";
import { ethers } from "ethers";
import { Presets } from "userop";
import { useUser } from "./context/UserContext";
import { UserOperationBuilder } from "userop";
import links from './links';
import { useState } from "react";




const SendUserOp = () => {
    
    const { user } = useUser();
    const { web3 } = useWeb3();
    const l = links();
    const [initialData, setInitialData] = useState<string>();
    const [unwrapData, setUnwrapData] = useState<string>();
    const [postData, setPostData] = useState<string>();
    const [signature, setSignature] = useState<string>();
    

    const handleSendUserOp = async () => {
        await fetch(l.initial + 'owner=' + user )
        .then((res) => {
            if(!res.ok){
                throw Error ('Failed to Fetch');
            }
            return(res.json());
        })
        .then(data => {
            console.log(data);
            setInitialData(data);
        })
        .catch((e) => {
            console.log(e.message);
        })
        
    }

    const handleUnwrap = async () => {
        await fetch(l.unwrap + 'target=2791Bca1f2de4661ED88A30C99A7a9449Aa84174&to=' + user + '&amount=100000000&owner=' + user)
        .then((res) => {
            if(!res.ok){
                throw Error ('Failed to Fetch');
            }
            return(res.json());
        })
        .then(data => {
            console.log(data);
            const stringified = JSON.stringify(data);
            setUnwrapData(stringified);
        })
        .catch((e) => {
            console.log(e.message);
        })
    }

    const handlePost = async () => {
        const obj = unwrapData? JSON.parse(unwrapData) : null;
        const nonce = obj.nonce;
        const sender = obj.sender;
        const callData = obj.callData;
        const callGasLimit = obj.callGasLimit;
        const initCode = obj.initCode;
        const maxFeePerGas = obj.maxFeePerGas;
        const maxPriorityFeePerGas = obj.maxPriorityFeePerGas;
        const paymasterAndData = obj.paymasterAndData;
        const entryPoint = '0x5FF137D4b0FDCD49DcA30c7CF57E578a026d2789';
        console.log(paymasterAndData);
        const preVerificationGas = obj.preVerificationGas;
        console.log(preVerificationGas);
        const verificationGasLimit = obj.verificationGasLimit;

        const signUserOperation = async (ctx : any) => {
            if(web3 && user){// Use the required signature scheme based on your wallet.
            // ctx.getRequestId() will generate the required hash for verification.
            // Multisig, ECDSA, etc.
                const hash = ctx.getUserOpHash().toString();
                const signable = hash.substr(2, hash.length);
                console.log('Hash: ' + hash);
                console.log('Signable: ' + signable);
                const signedMessage = await web3.eth.personal.sign(signable, user, "");
                setSignature(signedMessage);
                console.log(signedMessage);
                ctx.op.signature = signedMessage;
            }
        };
        
        const builder = new UserOperationBuilder()
        .setNonce(nonce)
        .setSender(sender)
        .setCallData(callData)
        .setCallGasLimit(callGasLimit)
        .setInitCode(initCode)
        .setMaxFeePerGas(maxFeePerGas)
        .setMaxPriorityFeePerGas(maxPriorityFeePerGas)
        .setPaymasterAndData(paymasterAndData)
        .setPreVerificationGas(preVerificationGas)
        .setVerificationGasLimit(verificationGasLimit)
        .useMiddleware(signUserOperation)
        console.log(builder);

        let op = await builder.buildOp(entryPoint , 137);
        console.log(op);
        
        const postData = { entryPoint, op };
        console.log(postData);
        

        // let userOp = await builder.getOp();

        // console.log(JSON.stringify(builder));
        // console.log(signature);


        await fetch(l.post, {
            method: 'POST',
            headers: {'Content-Type': 'application/json', 'Accept': 'application/json' },
            body: JSON.stringify(postData)

        })
        .then((res) => {
            if(!res.ok){
                throw Error ('Failed to Fetch');
            }
            return(res.json());
        })
        .then(data => {
            console.log(data);
            const stringified = JSON.stringify(data);
            setUnwrapData(stringified);
        })
        .catch((e) => {
            console.log(e.message);
        })
    }

    return (
        <div className="userOpButtons">
            {!initialData &&
                <Button onClick={handleSendUserOp}>SendUserOp</Button>
            }
            <p className="data">{initialData}</p>
            {initialData && !unwrapData &&
                <div>
                    <Button onClick={handleUnwrap}>Unwrap</Button>
                </div>
            }
            <p className="data">{unwrapData}</p>
            {unwrapData && initialData && 
                <div>
                    <Button onClick={handlePost}>Post </Button>
                </div>
            }
        </div>
    )
}
 
export default SendUserOp;