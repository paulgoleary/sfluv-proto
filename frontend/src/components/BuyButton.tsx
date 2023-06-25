import React, {useCallback, useEffect, useState} from 'react';
import {RatioButton} from "@ratio.me/ratiokit-react";
import {configureChains, mainnet, useAccount, useNetwork, useSignMessage} from 'wagmi';
import { publicProvider } from 'wagmi/providers/public'
import ConnectButton from "./ConnectButton";

const { chains, publicClient, webSocketPublicClient } = configureChains(
    [mainnet],
    [publicProvider()],
)

const BuyButton = () => {
    const {address, isConnected} = useAccount()
    const {chain} = useNetwork()
    const {signMessageAsync} = useSignMessage()

    const fetchSessionToken = useCallback(async (): Promise<string | null> => {
        try {
            const requestHeaders: HeadersInit = new Headers();
            requestHeaders.set('Content-Type', 'application/json');
            let sessionTokenResponse = await fetch(
                'http://localhost:8080/ratio/client/sessions',
                {
                    method: 'POST',
                    headers: requestHeaders,
                    body: JSON.stringify({
                        signingAddress: address,
                        depositAddress: address,
                        signingNetwork: chain?.name.toUpperCase(),
                    }),
                }
            );

            let data = await sessionTokenResponse.json();
            return data.id;
        } catch (e) {
            console.error(e);
        }
        return null;
    },[address, isConnected, chain])

    return <RatioButton
        text={'Buy with Ratio'}
        redirectUri={'https://yoursite.com/plaid/oauth'}
        fetchSessionToken={async () => {
            if(isConnected){
                return await fetchSessionToken()
            }
            return null
        }}
        signingCallback={async (challenge:string) => {
            return await signMessageAsync({
                message: challenge,
            })
        }}
    />
}

export default BuyButton