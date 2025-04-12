import Link from "next/link";
import styles from "./navbar.module.css";
import {links} from "./data";



import ButtonLogout from "@/elements/buttonLogout/button";
import Logo from "@/elements/logo/logo";



export default function Navbar(){
    return(
        
        <div className={styles.container}>
            <Logo />
            <div className={styles.links}>
                {links.map(link => <Link key={link.id} href={link.url} className={styles.link}>{link.title}</Link>)}
                <ButtonLogout />
            </div>
            
        </div>
    )
}