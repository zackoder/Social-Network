"use client";

import { useState, useEffect } from "react";
import styles from "./buttonFollow.module.css";
import { isAuthenticated } from "@/app/page";
import { FaUserPlus, FaUserCheck, FaUserClock } from "react-icons/fa";

export default function ButtonFollow({ profileId }) {
  const [isFollowing, setIsFollowing] = useState(false);
  const [isPending, setIsPending] = useState(false);
  const [currentUserId, setCurrentUserId] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const host = process.env.NEXT_PUBLIC_HOST
  // Fetch current user data when component mounts
  useEffect(() => {
    const fetchCurrentUser = async () => {
      try {
        const response = await fetch(`${host}/userData`,{
            credentials: "include",
          }
        );

        if (!response.ok) {
        }

        const data = await response.json();
        setCurrentUserId(data.Id || data.id);

        // Now check follow status
        if (profileId && (data.Id || data.id)) {
          checkFollowStatus(data.Id || data.id, profileId);
        }
      } catch (err) {
        // console.error("Error fetching current user:", err);
        isAuthenticated(response.status,"you should login first")
      } finally {
        setIsLoading(false);
      }
    };

    fetchCurrentUser();
  }, [profileId]);

  // Check if the current user is following the profile
  const checkFollowStatus = async (follower, followed) => {
    if (!follower || !followed) return;

    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_HOST}/api/checkFollowStatus?follower=${follower}&followed=${followed}`,
        {
          credentials: "include",
        }
      );

      if (response.ok) {
        const data = await response.json();
        setIsFollowing(data.isFollowing === true);
        setIsPending(data.isPending === true);
      }
    } catch (err) {
      // console.error("Error checking follow status:", err);
    }
  };

  const handleFollowToggle = async () => {
    if (!currentUserId || !profileId || isLoading) return;

    // Don't allow following yourself
    if (currentUserId.toString() === profileId.toString()) {
      setError("You cannot follow yourself");
      return;
    }

    try {
      const response = await fetch(`${host}/followReq?follower=${currentUserId}&followed=${profileId}`,
        {
          method: "POST",
          credentials: "include",
        }
      );

      // Only try to parse JSON if content-type is application/json
      const contentType = response.headers.get("content-type");
      let responseData;

      if (contentType && contentType.includes("application/json")) {
        responseData = await response.json();
      } else {
        // Just get the text for error logging
        responseData = await response.text();
      }

      if (!response.ok) {
        // console.error(
        //   `Failed to update follow status: ${response.status} ${response.statusText}`,
        //   responseData
        // );
    
      }

      // Handle the response
      if (typeof responseData === "object") {
        if (responseData.resp === "followed seccessfoly") {
          setIsFollowing(true);
          setIsPending(false);
        } else if (responseData.resp === "unfollowed seccessfoly") {
          setIsFollowing(false);
          setIsPending(false);
        } else if (responseData.resp === "follow request sent") {
          setIsPending(true);
        }
      } else {
        // The response wasn't JSON, toggle the state based on previous state
        setIsFollowing(!isFollowing);
      }
    } catch (err) {
      // console.error("Error toggling follow status:", err);
        isAuthenticated(response.status, "you should login first")
    
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
      className={`${styles.button} ${isFollowing ? styles.following : ""}`}
      onClick={handleFollowToggle}
      disabled={isLoading}
    >
      {isPending ? (
        <>
          <FaUserClock /> Pending
        </>
      ) : isFollowing ? (
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
