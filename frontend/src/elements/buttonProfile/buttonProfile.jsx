"use client";

import Link from "next/link";
import Image from "next/image";
import styles from "./buttonProfile.modules.css";
import { useState, useEffect } from "react";

export default function ButtonProfile() {
  const [userData, setUserData] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await fetch(
          `${process.env.NEXT_PUBLIC_HOST}/userData`,
          {
            credentials: "include",
          }
        );
            console.log("response");
            
        if (!response.ok) {
          throw new Error("Failed to fetch user data");
        }

        const data = await response.json();
        console.log("User data received:", data);

        setUserData(data);
      } catch (err) {
        console.error("Error fetching user data:", err);
        setError(err.message);
      } finally {
        setIsLoading(false);
      }
    };

    fetchUserData();
  }, []);

  // Change the URL to use the existing profile route insyprofile
  const profileUrl = `/Myprofile` ;

  return (
    <Link href={profileUrl}>
      <Image
        src="/images/profile.png"
        className="imgProfile"
        width={38}
        height={38}
        alt="profile"
        title="profile"
      />
    </Link>
  );
}
