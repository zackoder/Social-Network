"use client";
import { FaCloudUploadAlt } from "react-icons/fa";
import { useState, useContext, useEffect } from "react";
import "./createPost.modules.css";
// import { getData } from "../post/post";
import { useRef } from "react";
import ContactsPrivate from "../contactprivate/contactprivate";
import { isAuthenticated } from "@/app/page";
import { DataContext } from "@/contexts/dataContext";
const host = process.env.NEXT_PUBLIC_HOST;
export default function CreatePost({ onPostCreated }) {
  const [userData, setUserData] = useState([]);
  useEffect(() => {
    try {
      const fetchUserData = async () => {
        const response = await fetch(`${host}/userData`, {
          credentials: "include",
        });
        if (!response.ok) {
          isAuthenticated(response.status, "you should login first");
        }
        const userdata = await response.json();

        setUserData(userdata);
      };
      fetchUserData();
    } catch (error) {
      console.log("we cant fetch user data for post", error);
    }
  }, []);

  let [privacy, setPrivacy] = useState("public");
  let [title, setTitle] = useState("");
  let [content, setContent] = useState("");
  let [image, setImage] = useState("");
  const { selectedContactsIds, setSelectedContactsIds } =
    useContext(DataContext);
  let friends = selectedContactsIds;
  const fileInputRef = useRef(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    // console.log("friends", friends);
    const formData = new FormData();
    const postData = {
      privacy,
      title,
      content,
      friends,
    };

    formData.append("postData", JSON.stringify(postData));
    if (image) {
      formData.append("avatar", image);
    }

    try {
      const response = await fetch(`${host}/addPost`, {
        method: "POST",
        body: formData,
        credentials: "include",
      });

      const newPost = await response.json();

      if (!response.ok) {
        isAuthenticated(response.status, newPost.error);
        console.log(newPost.error);
        return;
      } else {
        // Reset form
        setPrivacy("public");
        setTitle("");
        setContent("");
        setImage("");
        setSelectedContactsIds([]);
      }

      // Reset file input
      if (fileInputRef.current) {
        fileInputRef.current.value = "";
      }

      // Notify parent
      if (onPostCreated) {
        onPostCreated(newPost);
      }
    } catch (error) {
      console.log("Submission error:", error);
    }
  };

  return (
    <div className="postContainer">
      <form onSubmit={handleSubmit}>
        <div className="identityProfile">
          <div className="imageProfile">
            <img
              className="imageProfile"
              src={`http://${userData.avatar}`}
              alt="profile"
            />
          </div>
          <div className="nameProfile">
            <div className="name-privacy">
              <h3>{userData.firstName}</h3>
              <select
                onChange={(e) => {
                  setPrivacy(e.target.value);
                }}
                name="friends"
                id="friends"
                defaultValue={"public"}
              >
                <option value={"public"}>Public</option>
                <option value={"private"}>Private</option>
                <option value={"almostPrivate"}>Almost Private </option>
              </select>
            </div>
            <div>
              {privacy === "private" && (
                <ContactsPrivate className="contactPrivate" />
              )}
            </div>
          </div>
        </div>
        <div className="title">
          <input
            onChange={(e) => {
              setTitle(e.target.value);
            }}
            value={title}
            type="text"
            name="title"
            placeholder="enter your title"
          />
        </div>
        <div className="content">
          {/* <textarea name="content" placeholder="enter your content"></textarea> */}
          <textarea
            onChange={(e) => {
              setContent(e.target.value);
            }}
            value={content}
            placeholder="Write a content..."
            // value={comment}
            // onChange={(e) => setComment(e.target.value)}
            rows={4}
            style={{
              width: "90%",
              padding: "5px",
                borderRadius: "4px",
                resize: "none",
                outline: "none",
                border: "none",
                marginLeft: "5%",
                backgroundColor: "#333",
              }}
            />
          </div>
          <div className="uploadImage">
            <input
              onChange={(e) => {
                setImage(e.target.files[0]);
              }}
              id="uploadImage"
              className="hiddenInput"
              ref={fileInputRef}
              type="file"
            />
            <label htmlFor="uploadImage" className="uploadLabel">
              <FaCloudUploadAlt className="iconUpload" />
            </label>
          </div>
          <div>
            <input className="submit" type="submit" value="Publish" />
          </div>
        </form>
      </div>
    );
  }
