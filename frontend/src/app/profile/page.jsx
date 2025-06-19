"use client";
 
import { useEffect, useState, useCallback } from "react";
import { useSearchParams } from "next/navigation";
import ButtonFollow from "@/elements/buttonFollow/buttonFollow";
import styles from "./profile.module.css";
import Image from "next/image";
import Post from "@/components/post/post";
import { FaLock, FaLockOpen } from "react-icons/fa";
import { isAuthenticated } from "../page";
import { debounce } from "@/utils/debounce";

export default function ProfilePage() {
  const [offset, setOffset] = useState(0);
  const [hasMore, setHasMore] = useState(true);
  const [loading, setLoading] = useState(false);
  const LIMIT = 10;

  const searchParams = useSearchParams();
  const profileId = searchParams.get("id");
  const [profile, setProfile] = useState(null);
  const [posts, setPosts] = useState([]);
  const [followers, setFollowers] = useState([]);
  const [following, setFollowing] = useState([]);
  const [isPrivate, setIsPrivate] = useState("");
  const [isLoading, setIsLoading] = useState(true);
  const [activeTab, setActiveTab] = useState("posts");
  const [showModal, setShowModal] = useState(false);
  const [modalContent, setModalContent] = useState({ title: "", data: [] });
  const host = process.env.NEXT_PUBLIC_HOST;
  const [error, setError] = useState("");

  console.log("get id of the user ", profileId);

  useEffect(() => {
    if (profileId) {
      fetchProfileData();
    }
  }, [profileId]);

  const debouncedFetchPosts = useCallback(debounce(fetchProfileposts, 300), [
    offset,
    hasMore,
    loading,
  ]);

  async function fetchProfileposts() {
    if (loading || !hasMore) return;
    setLoading(true);
    try {
      const response = await fetch(
        `${host}/api/getProfilePosts?id=${profileId}&offset=${offset}&limit=${LIMIT}`,
        {
          // cache: "no-store",
          method: "GET",
          credentials: "include",
        }
      );
      if (!response.ok) {
        // throw new Error("Failed to fetch posts");
      }
      const data = await response.json();
      if (data !== null && data.message === "this profile is private") {
        setPosts([]);
        return;
      }  
      console.log('befoore entring the condition' ,hasMore);
      if (posts.length <= offset &&data === null ) {
        console.log('im here in the condition',hasMore);
        setHasMore(false); // No more posts available
        return;
      }
      setPosts((prev) => [...prev, ...data]);
      setOffset((prev) => prev + LIMIT);
    } catch (err) {
      console.error("Fetch error:", err);
    } finally {
      console.log(loading, hasMore);
      // setHasMore(true);
      setLoading(false);
    }
  }

  const fetchProfileData = async () => {
    setIsLoading(true);
    await fetchProfileposts();
    try {
      // Fetch profile information
      const profileResponse = await fetch(
        `${host}/api/registrationData?id=${profileId}`,
        { credentials: "include" }
      );
      const profileData = await profileResponse.json();
      console.log("wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww",profileData);
      

      if (!profileResponse.ok) {
        console.log(`Profile response error: ${profileResponse.status}`);
        isAuthenticated(profileResponse.status, profileData.error);
      }

      setProfile(profileData.registration_data);
      setIsPrivate(profileData.profile_status);

      // const postsResponse = await fetch(
      //   `${host}/api/getProfilePosts?id=${profileId}`,
      //   {
      //     credentials: "include",
      //   }
      // );
      // const posts = await postsResponse.json();

      //  console.log("this is for posts", posts);

      // Ensure posts is always an array
      // setPosts(Array.isArray(posts) ? posts : []);
      // console.error("Failed to fetch posts");
      // setPosts([]); // Set empty array on error

      // Fetch followers
      const followers = await fetch(
        `${host}/api/getfollowers?id=${profileId}`,
        {
          credentials: "include",
        }
      );
      const followersData = await followers.json();

      if (followers.ok) {
        setFollowers(Array.isArray(followersData) ? followersData : []);
      }
      //  console.log("followersResponse.ok", followers);

      // Fetch following
      const following = await fetch(
        `${host}/api/getfollowinglist?id=${profileId}`,
        {
          credentials: "include",
        }
      );

      if (following.ok) {
        const followingData = await following.json();
        setFollowing(Array.isArray(followingData) ? followingData : []);
      }
    
      // console.log("followers.ok", following);
    } catch (error) {
      // console.error("Error in profile data fetch:", error);
      setError(error.message || "Error loading profile");
    } finally {
      setIsLoading(false);
    }
    
    // if (isPrivate === "private" || isPrivate === "public") {
    //   await handlePrivacyToggle();
    // }
  };
  

  const handlePrivacyToggle = async () => {
    try {
      const response = await fetch(`${host}/updatePrivacy`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ isPrivate: isPrivate }),
      });
      const data = await response.json();
      // console.log(";;;;;;;;;;;;;;;;;;;;;;;;;;", data.profile_status);

      if (response.ok) {
        setIsPrivate(data.profile_status);
      }
    } catch (error) {
      console.log("Error updating privacy settings", error);
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
            <img
              className={styles.image}
              src={`http://${profile.avatar}`}
              alt={`${profile.firstName} ${profile.lastName}`}
              // fill={true}
              />
          </div>
          <div className={styles.name}>
            <h3>
              {profile.firstName} {profile.lastName}
            </h3>
            {profile.nickName && <p>@{profile.nickName}</p>}
            <p>{profile.aboutMe}</p>

            {["public", "private"].includes(isPrivate) ? (
              <div className={styles.privacyToggle}>
                <span>
                  {isPrivate === "private" ? <FaLock /> : <FaLockOpen />}
                </span>
                <label className={styles.toggleSwitch}>
                  <input
                    type="checkbox"
                    checked={isPrivate === "private"}
                    onChange={handlePrivacyToggle}
                  />
                  <span className={styles.slider}></span>
                </label>
                <span>
                  {isPrivate === "private"
                    ? "Private Profile"
                    : "Public Profile"}
                </span>
              </div>
            ) : (
              <ButtonFollow profileId={profileId} />
            )}
          </div>
        </div>

        <div className={styles.follow}>
          <div className={styles.followers} onClick={showFollowers}>
            <p>Followers</p>
            <h3>{followers ? followers.length : "0"}</h3>
          </div>
          <div className={styles.following} onClick={showFollowing}>
            <p>Following</p>
            <h3>{following ? following.length : "0"}</h3>
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

        {activeTab === "posts" &&
          (posts && posts.length > 0 ? (
            <>
                <Post posts={posts} />
              {hasMore ?(
                <button onClick={debouncedFetchPosts} disabled={loading}>
                  {loading ? "Loading..." : "Load More"}
                </button>
              ) : (
                <button>{"there are no more posts"}</button>
              )}
            </>
          ) : (
            <div className={styles.noContent}>No posts to display</div>
          ))}
        {/* {activeTab === "posts" &&
          (posts && posts.length > 0 ? (
            <>
              <Post posts={posts} />
              {hasMore && (
                <button onClick={debouncedFetchPosts} disabled={loading}>
                  {loading ? "Loading..." : "Load More"}
                </button>
              )}
            </>
          ) : (
            <div className={styles.noContent}>No posts to display</div>
          ))} */}
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
                    <img
                      className={styles.userAvatar}                      
                      src={`${host}${user.avatar}`}
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
