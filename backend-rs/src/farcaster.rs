

#[cfg(test)]
mod tests {

    use ethers::{
        prelude::{abigen, Abigen},
        providers::{Http, Provider},
        types::Address,
    };
    use eyre::Result;
    use std::sync::Arc;
    use ethers::abi::AbiEncode;
    use ethers::prelude::U256;
    use tokio_test::*;

    #[tokio::test]
    async fn test_id_registry() {

        abigen!(IdRegistry, "./src/abi/IdRegistry.json");

        const RPC_URL: &str = "https://opt-mainnet.g.alchemy.com/v2/4BTkAMn0RXvlAsLiniNtAdVNDH3PYcdE"; // TODO: env !!!
        const ID_REGISTRY_ADDRESS: &str = "0x00000000fc6c5f01fc30151999387bb99a9f489b";
        const TEST_ADDRESS: &str = "0x8531eccb1a475f09229bd986abe8b69c1d17ac04";

        let provider = Provider::<Http>::try_from(RPC_URL).unwrap();
        let client = Arc::new(provider);
        let address: Address = ID_REGISTRY_ADDRESS.parse().unwrap();
        let contract = IdRegistry::new(address, client);

        let test_address: Address = TEST_ADDRESS.parse().unwrap();

        let id = contract.id_of(test_address).await.unwrap();
        assert_eq!(id.as_u64(), 191617);

        let recovery = contract.recovery_of(191617.into()).await.unwrap();
        println!("{}", recovery.encode_hex());

        contract.tru
    }
}