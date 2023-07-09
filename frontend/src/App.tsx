import { VStack } from "@chakra-ui/react"
import { useUser } from "./context/UserContext"
import ConnectButton from "./components/ConnectButton"
import WalletDetail from "./components/WalletDetail"
import DisconnectButton from "./components/DisconnectButton"
import ShowUIButton from "./components/ShowUIButton"
import SignMessage from "./components/SignMessage"
import BuyButton from "./components/BuyButton";
import Profile from "./Profile"
import Navbar from "./Navbar"
import Home from "./Home"
import About from "./About"
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';


function App() {

  return (
      <Router>
        <div className="App">
          <Navbar/>
            <div className="content">
              <Routes>
                <Route path="/profile" element={<Profile/>}/>
                <Route path="/" element={<Home/>}/>
                <Route path="/about" element={<About/>}/>
              </Routes>
            </div>
        </div>
      </Router>
  )
}

export default App
