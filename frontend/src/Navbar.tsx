import { Link } from "react-router-dom";
import logo from"./imgs/SFLUV Banner.png";
import profileImg from "./imgs/Profile.png";

const Navbar = () => {
    return (
        <nav className="navbar">
            <Link to="/"><img src={logo} alt="SFLUV" className="logo"/></Link>
            <div className="web-links">
                <Link to="/about">About</Link>
            </div>
            <div className="app-links">
                <Link to="/profile"><img src={profileImg} alt="Profile" className="profileImg"/></Link>
            </div>
        </nav>
      );
}
 
export default Navbar;