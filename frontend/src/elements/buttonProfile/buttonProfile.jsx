import Link from "next/link"
import Image from "next/image";
import styles from "./buttonProfile.modules.css"

export default function ButtonProfile(){
    return (
        <Link href="/profile/2">
            <Image 
                src="/images/profile.png"
                // className={styles.imgProfile}
                className="imgProfile"
                width={38}
                height={38}
                alt="profile"
                title="profile"
            />
        </Link>
    );
}