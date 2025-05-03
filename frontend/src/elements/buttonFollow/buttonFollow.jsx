'use client'
import { useState, useEffect } from "react";
import styles from "./buttonFollow.module.css";
import { RiUserFollowFill, RiUserUnfollowFill } from "react-icons/ri";

export default function ButtonFollow({ profileId }) {
    const [isFollowing, setIsFollowing] = useState(false);
    const [isLoading, setIsLoading] = useState(false);

    // Check if the current user is following the profile owner
    useEffect(() => {
        if (profileId) {
            checkFollowStatus();
        }
    }, [profileId]);

    const checkFollowStatus = async () => {
        try {
            // Call API to check if user is following this profile
            const response = await fetch(`${process.env.NEXT_PUBLIC_HOST}/api/checkFollow?profileId=${profileId}`, {
                method: 'GET',
                credentials: 'include',
            });
            
            if (response.ok) {
                const data = await response.json();
                setIsFollowing(data.isFollowing);
            }
        } catch (error) {
            console.error("Failed to check follow status:", error);
        }
    };

    const handleFollowToggle = async () => {
        if (isLoading) return;
        
        setIsLoading(true);
        try {
            const endpoint = isFollowing ? '/api/unfollow' : '/api/follow';
            const response = await fetch(`${process.env.NEXT_PUBLIC_HOST}${endpoint}?followed=${profileId}`, {
                method: 'POST',
                credentials: 'include',
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            if (response.ok) {
                setIsFollowing(!isFollowing);
            } else {
                console.error("Failed to update follow status");
            }
        } catch (error) {
            console.error("Error updating follow status:", error);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <button 
            className={`${styles.button} ${isFollowing ? styles.following : ''}`} 
            onClick={handleFollowToggle}
            disabled={isLoading}
        >
            {isFollowing ? (
                <>
                    <RiUserUnfollowFill /> Unfollow
                </>
            ) : (
                <>
                    <RiUserFollowFill /> Follow
                </>
            )}
        </button>
    );
}