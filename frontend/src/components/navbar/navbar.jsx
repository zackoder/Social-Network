import styles from "./navbar.module.css";
import Notification from "../notifications/notification";



import ButtonLogout from "@/elements/buttonLogout/button";
import Logo from "@/elements/logo/logo";



export default function Navbar(){
    return(
        <div className={styles.container}>
            <Logo />
            <Notification />
            <ButtonLogout />
        </div>
    )
}