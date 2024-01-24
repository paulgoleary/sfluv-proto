const getPrivateKey = async (web3auth) => {
  const privateKey = await web3auth.provider.request({
    method: "eth_private_key"
  });
  return(privateKey);
}

export default getPrivateKey;