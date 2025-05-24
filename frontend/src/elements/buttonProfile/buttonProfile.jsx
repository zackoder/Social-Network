"use client";

import Image from "next/image";
import styles from "./buttonProfile.modules.css";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";

export default function ButtonProfile() {
  const [userData, setUserData] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const router = useRouter();
  
  useEffect(() => {
    const fetchUserData = async () => {
      try {
        console.log(
          "Fetching user data from:",
          `${process.env.NEXT_PUBLIC_HOST}/userData`
        );
        const response = await fetch(
          `${process.env.NEXT_PUBLIC_HOST}/userData`,
          {
            credentials: "include",
          }
        );
        console.log("response the button profile", response);
        
        if (!response.ok) {
          // If response is 401 Unauthorized, redirect to login
          if (response.status === 401) {
            router.push("/login");
            return;
          }
          // throw new Error(`Failed to fetch user data: ${response.status}`);
        }

        // Handle different response content types
        const contentType = response.headers.get("content-type");
        if (!contentType || !contentType.includes("application/json")) {
          console.log("Non-JSON response received:", contentType);
          const text = await response.text();
          console.log("Response text:", text);
          // throw new Error("Received non-JSON response from server");
        }

        const data = await response.json();
        console.log("User data received:", data);

        if (!data || Object.keys(data).length === 0) {
          console.log("Empty user data received");
          // throw new Error("Invalid user data");
        }

        // Check multiple possible ID field names
        const userId =
          data.Id ||
          data.id ||
          data.ID ||
          data.userId ||
          data.user_id;
        if (!userId) {
          console.log("User data does not contain ID field:", data);
          // throw new Error("User ID missing from response");
        }

        setUserData(data);
      } catch (err) {
        console.error("Error fetching user data:", err);
        setError(err.message);
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
      console.log("Cannot navigate due to error:", error);
      router.push("/login");
      return;
    }

    // Look for ID in different possible case formats
    const userId =
      userData?.Id ||
      userData?.id ||
      userData?.ID ||
      userData?.userId ||
      userData?.user_id;

    if (userId) {
      console.log(`Navigating to profile with ID: ${userId}`);
      router.push(`/profile?id=${userId}`);
    } else {
      console.log("Cannot navigate: user ID not found in data:", userData);
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
