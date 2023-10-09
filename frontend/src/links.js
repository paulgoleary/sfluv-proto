const links = () => {
  const links = {
      initial: 'http://localhost:8080/erc4337/sender?',
      wrap: 'http://localhost:8080/erc4337/userop/approve?',
      unwrap: 'http://localhost:8080/erc4337/userop/withdrawto?',
      post: 'http://localhost:8080/erc4337/userop/send'
  };
  return(links);
};

export default links;