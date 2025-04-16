import styles from "./navbar.module.css";
import Notification from "../notifications/notification";



import ButtonLogout from "@/elements/buttonLogout/button";
import Logo from "@/elements/logo/logo";



export default function Navbar(){
    return(
        <div className={styles.container}>
            <div className={styles.nav_start}>
                <Logo />
            </div>
            <div className={styles.nav_middle}>
                <Notification />
            </div>
            <div className={styles.nav_end}>
                <ButtonLogout />
            </div>
        </div>
    )
}