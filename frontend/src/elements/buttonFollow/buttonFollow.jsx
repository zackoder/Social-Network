"use client";

import { useState, useEffect } from "react";
import styles from "./buttonFollow.module.css";
import { isAuthenticated } from "@/app/page";
import { FaUserPlus, FaUserCheck, FaUserClock } from "react-icons/fa";

export default function ButtonFollow({ profileId }) {
  const [followstatus, setfollowststus] = useState("");
  // const [isFollowing, setIsFollowing] = useState(false);
  // const [isPending, setIsPending] = useState(false);
  const [currentUserId, setCurrentUserId] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const host = process.env.NEXT_PUBLIC_HOST;
  // Fetch current user data when component mounts
  useEffect(() => {
    const fetchCurrentUser = async () => {
      try {
        const response = await fetch(`${host}/userData`, {
          credentials: "include",
        });
        const data = await response.json();
        // console.log("jjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjjj", data);

        setCurrentUserId(data.id);
        // Now check follow status
        if (profileId && data.id) {
          checkFollowStatus(data.id, profileId);
        }
      } catch (err) {
        console.error("Error fetching current user:", err);
        isAuthenticated(response.status, "you should login first");
      } finally {
        setIsLoading(false);
      }
    };

    fetchCurrentUser();
  }, [profileId]);

  // Check if the current user is following the prof   setIsFollowing(data.isFollowing === true);
  // setIsPending(data.isPending === true);ile
  const checkFollowStatus = async (follower, followed) => {
    if (!follower || !followed) return;

    try {
      const response = await fetch(
        `${host}/api/registrationData?id=${profileId}`,
        {
          credentials: "include",
        }
      );

      if (response.ok) {
        const data = await response.json();
        console.log(data);
        setfollowststus(data.profile_status);
      }
    } catch (err) {
      console.error("Error checking follow status:", err);
    }
  };

  const handleFollowToggle = async () => {
    // if (!currentUserId || !profileId || isLoading) return;

    // Don't allow following yourself
    if (currentUserId.toString() === profileId.toString()) {
      setError("You cannot follow yourself");
      return;
    }

    try {
      const response = await fetch(`${host}/followReq?followed=${profileId}`, {
        method: "POST",
        credentials: "include",
      });

      const data = await response.json();

      if (data.resp === "followed seccessfoly") {
        setfollowststus("unfollow");
      } else if (data.resp === "unfollowed seccessfoly") {
        setfollowststus("follow");
      } else if (data.resp === "follow request sent") {
        setfollowststus("follow sent");
      }
    } catch (err) {
      console.error("Error toggling follow status:", err);
      isAuthenticated(response.status, "you should login first");
    }
  };

  // Don't show button when viewing your own profile or during initial load
  if (currentUserId && currentUserId.toString() === profileId?.toString()) {
    return null;
  }
  
  // Don't show button while loading
  if (isLoading) {
    return (
      <button className={styles.button} disabled>
        Loading...
      </button>
    );
  }

  return (
    <button
      className={`${styles.button} ${
        followstatus == "follow" ? styles.following : ""
      }`}
      onClick={handleFollowToggle}
    >
      {followstatus == "follow sent" ? (
        <>
          <FaUserClock /> follow sent
        </>
      ) : followstatus == "unfollow" ? (
        <>
          <FaUserCheck /> Following
        </>
      ) : (
        <>
          <FaUserPlus /> Follow
        </>
      )}
    </button>
  );
}
