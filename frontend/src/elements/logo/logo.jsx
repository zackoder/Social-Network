import Link from "next/link";
import styles from "./logo.module.css"

import {Montserrat} from "next/font/google";
const fontLogo = Montserrat({
    subsets: ['latin'],
    weight: ['700']
});

export default function Logo(){
    return(
        <Link href="/" className={`${styles.logo} ${fontLogo.className}`}>
            Social-Network
        </Link>
    );
}