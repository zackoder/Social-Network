"use client";

import { useEffect, useState } from "react";
import { useSearchParams } from "next/navigation";
import ButtonFollow from "@/elements/buttonFollow/buttonFollow";
import styles from "./profile.module.css";
import Image from "next/image";
import Post from "@/components/post/post";
import { FaLock, FaLockOpen } from "react-icons/fa";

export default function ProfilePage() {
  const searchParams = useSearchParams();
  const profileId = searchParams.get("id");

  const [profile, setProfile] = useState(null);
  const [posts, setPosts] = useState([]);
  const [followers, setFollowers] = useState([]);
  const [following, setFollowing] = useState([]);
  const [isOwnProfile, setIsOwnProfile] = useState(false);
  const [isPrivate, setIsPrivate] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const [activeTab, setActiveTab] = useState("posts");
  const [showModal, setShowModal] = useState(false);
  const [modalContent, setModalContent] = useState({ title: "", data: [] });

  useEffect(() => {
    if (profileId) {
      fetchProfileData();
    }
  }, [profileId]);

  const fetchProfileData = async () => {
    setIsLoading(true);
    try {
      // Fetch profile information
      console.log(
        `${process.env.NEXT_PUBLIC_HOST}/api/getProfilePosts?id=${profileId}`
      );
      const profileResponse = await fetch(
        `${process.env.NEXT_PUBLIC_HOST}/api/registrationData?id=${profileId}`,
        {
          credentials: "include",
        }
      );
      let dataxx = await profileResponse.json();
      console.log(dataxx);
      if (profileResponse.ok) {
        const profileData = await profileResponse.json();
        setProfile(profileData);
        setIsPrivate(profileData.privacy === "private");
        setIsOwnProfile(profileData.isOwnProfile);
      } else {
        console.error("Failed to fetch profile data");
      }
      // Fetch posts
      const postsResponse = await fetch(
        `${process.env.NEXT_PUBLIC_HOST}/api/getProfilePosts?id=${profileId}`,
        {
          credentials: "include",
        }
      );

      if (postsResponse.ok) {
        const postsData = await postsResponse.json();
        setPosts(postsData);
      } else {
        console.error("Failed to fetch posts");
      }

      // Fetch followers
      const followersResponse = await fetch(
        `${process.env.NEXT_PUBLIC_HOST}/api/getfollowers?id=${profileId}`,
        {
          credentials: "include",
        }
      );

      if (followersResponse.ok) {
        const followersData = await followersResponse.json();
        setFollowers(followersData);
      } else {
        console.error("Failed to fetch followers");
      }

      // Fetch following
      const followingResponse = await fetch(
        `${process.env.NEXT_PUBLIC_HOST}/api/followers?id=${profileId}`,
        {
          credentials: "include",
        }
      );

      if (followingResponse.ok) {
        const followingData = await followingResponse.json();
        setFollowing(followingData);
      } else {
        console.error("Failed to fetch following");
      }
    } catch (error) {
      console.error("Error fetching profile data:", error);
    } finally {
      setIsLoading(false);
    }
  };

  const handlePrivacyToggle = async () => {
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_HOST}/api/togglePrivacy`,
        {
          method: "POST",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ isPrivate: !isPrivate }),
        }
      );

      if (response.ok) {
        setIsPrivate(!isPrivate);
      } else {
        console.error("Failed to update privacy settings");
      }
    } catch (error) {
      console.error("Error updating privacy settings:", error);
    }
  };

  const showFollowers = () => {
    setModalContent({
      title: "Followers",
      data: followers,
    });
    setShowModal(true);
  };

  const showFollowing = () => {
    setModalContent({
      title: "Following",
      data: following,
    });
    setShowModal(true);
  };

  if (isLoading) {
    return <div className={styles.container}>Loading...</div>;
  }

  if (!profile) {
    return <div className={styles.container}>Profile not found</div>;
  }

  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <div className={styles.info}>
          <div className={styles.boxImage}>
            <Image
              className={styles.image}
              src={profile.avatar || "/profile/profile.png"}
              alt={`${profile.firstName} ${profile.lastName}`}
              fill={true}
            />
          </div>
          <div className={styles.name}>
            <h3>
              {profile.firstName} {profile.lastName}
            </h3>
            {profile.nickName && <p>@{profile.nickName}</p>}
            <p>{profile.aboutMe}</p>

            {isOwnProfile ? (
              <div className={styles.privacyToggle}>
                <span>{isPrivate ? <FaLock /> : <FaLockOpen />}</span>
                <label className={styles.toggleSwitch}>
                  <input
                    type="checkbox"
                    checked={isPrivate}
                    onChange={handlePrivacyToggle}
                  />
                  <span className={styles.slider}></span>
                </label>
                <span>{isPrivate ? "Private Profile" : "Public Profile"}</span>
              </div>
            ) : (
              <ButtonFollow profileId={profileId} />
            )}
          </div>
        </div>

        <div className={styles.follow}>
          <div className={styles.followers} onClick={showFollowers}>
            <p>Followers</p>
            <h3>{followers.length}</h3>
          </div>
          <div className={styles.following} onClick={showFollowing}>
            <p>Following</p>
            <h3>{following.length}</h3>
          </div>
        </div>
      </header>

      <div className={styles.tabs}>
        <div
          className={`${styles.tab} ${
            activeTab === "posts" ? styles.active : ""
          }`}
          onClick={() => setActiveTab("posts")}
        >
          Posts
        </div>
        <div
          className={`${styles.tab} ${
            activeTab === "about" ? styles.active : ""
          }`}
          onClick={() => setActiveTab("about")}
        >
          About
        </div>
      </div>

      <main>
        {activeTab === "posts" && (
          <div>
            {posts.length > 0 ? (
              posts.map((post) => <Post key={post.id} post={post} />)
            ) : (
              <div className={styles.noContent}>No posts to display</div>
            )}
          </div>
        )}

        {activeTab === "about" && (
          <div className={styles.aboutUser}>
            <h4>User Information</h4>

            <div className={styles.infoRow}>
              <div className={styles.infoLabel}>Name:</div>
              <div>
                {profile.firstName} {profile.lastName}
              </div>
            </div>

            {profile.nickName && (
              <div className={styles.infoRow}>
                <div className={styles.infoLabel}>Nickname:</div>
                <div>{profile.nickName}</div>
              </div>
            )}

            <div className={styles.infoRow}>
              <div className={styles.infoLabel}>Age:</div>
              <div>{profile.age}</div>
            </div>

            <div className={styles.infoRow}>
              <div className={styles.infoLabel}>Gender:</div>
              <div>{profile.gender}</div>
            </div>

            <div className={styles.infoRow}>
              <div className={styles.infoLabel}>About:</div>
              <div>{profile.aboutMe || "No bio provided"}</div>
            </div>
          </div>
        )}
      </main>

      {showModal && (
        <div className={styles.modal} onClick={() => setShowModal(false)}>
          <div
            className={styles.modalContent}
            onClick={(e) => e.stopPropagation()}
          >
            <button
              className={styles.closeButton}
              onClick={() => setShowModal(false)}
            >
              ×
            </button>
            <h3>{modalContent.title}</h3>

            {modalContent.data.length > 0 ? (
              <ul className={styles.usersList}>
                {modalContent.data.map((user) => (
                  <li key={user.id} className={styles.userItem}>
                    <Image
                      className={styles.userAvatar}
                      src={user.avatar || "/profile/profile.png"}
                      alt={`${user.firstName} ${user.lastName}`}
                      width={40}
                      height={40}
                    />
                    <div className={styles.userName}>
                      {user.firstName} {user.lastName}
                    </div>
                    <ButtonFollow profileId={user.id} />
                  </li>
                ))}
              </ul>
            ) : (
              <div className={styles.noContent}>No users to display</div>
            )}
          </div>
        </div>
      )}
    </div>
  );
}
