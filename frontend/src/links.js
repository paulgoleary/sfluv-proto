const links = () => {
    // const base_url = 'http://localhost:8080'
    const base_url = 'http://ec2-54-219-196-236.us-west-1.compute.amazonaws.com:8080'
    const links = {
      initial: base_url + '/erc4337/sender?',
      wrap: base_url + '/erc4337/userop/approve?',
      unwrap: base_url + '/erc4337/userop/withdrawto?',
      post: base_url + '/erc4337/userop/send'
  };
  return(links);
};

export default links;