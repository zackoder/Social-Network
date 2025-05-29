"use client";

import Image from "next/image";
import styles from "./buttonProfile.modules.css";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { isAuthenticated } from "@/app/page";

export default function ButtonProfile() {
  const [userData, setUserData] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const router = useRouter();
  const host = process.env.NEXT_PUBLIC_HOST
  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const response = await fetch(
          `${host}/userData`,
          {
            credentials: "include",
          }
        );
        const data = await response.json();

        if (!response.ok) {
          // If response is 401 Unauthorized, redirect to login
          isAuthenticated(response.status, data.error)
        }
        
        // Handle different response content types
      

        console.log("User data received:", data);

        // if (!data || Object.keys(data).length === 0) {
        //   console.error("Empty user data received");
        //   throw new Error("Invalid user data");
        // }

        // Check multiple possible ID field names
        const userId = data.id;
        console.log(userId);
        
        localStorage.setItem("user-id", userId);
        if (!userId) {
          console.warn("User data does not contain ID field:", data);
        }

        setUserData(data);
      } catch (err) {

      } finally {
        setIsLoading(false);
      }
    };

    fetchUserData();
  }, [router]);

  const handleProfileClick = (e) => {
    e.preventDefault();

    if (isLoading) {
      console.log("Still loading user data, waiting...");
      return;
    }

    if (error) {
      // console.error("Cannot navigate due to error:", error);
      router.push("/login");
      return;
    }

    // Look for ID in different possible case formats
    const userId = userData.id;

    if (userId) {
      console.log(`Navigating to profile with ID: ${userId}`);
      router.push(`/profile?id=${userId}`);
    } else {
      // console.error("Cannot navigate: user ID not found in data:", userData);
      router.push("/login");
    }
  };

  return (
    <a
      href="#"
      onClick={handleProfileClick}
      style={{ cursor: isLoading ? "wait" : "pointer" }}
      aria-label="My Profile"
    >
      <Image
        src="/images/profile.png"
        className="imgProfile"
        width={38}
        height={38}
        alt="profile"
        title={
          isLoading
            ? "Loading profile..."
            : error
            ? "Error loading profile"
            : "My Profile"
        }
      />
    </a>
  );
}
