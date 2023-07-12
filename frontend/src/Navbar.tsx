import { Link } from "react-router-dom";
import logo from"./imgs/SFLUV Banner.png";
import profileImg from "./imgs/Profile.png";
import { useUser } from "./context/UserContext";

const Navbar = () => {
    const { user } = useUser();
    return (
        <nav className="navbar">
            <Link to="/"><img src={logo} alt="SFLUV" className="logo"/></Link>
            <div className="web-links">
                <Link to="/about">About</Link>
                
            </div>
            <div className="app-links">
                <Link to="/test">Test</Link>
                {user && <Link to="/buy">Buy SFLUV</Link>}
                <Link to="/profile"><img src={profileImg} alt="Profile" className="profileImg"/></Link>  
            </div>
        </nav>
      );
}
 
export default Navbar;