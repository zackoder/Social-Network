"use client"
import { FaCloudUploadAlt } from "react-icons/fa";
import { use, useState } from "react";
import "./createPost.modules.css"
import { getData } from "../post/post";
import { useRef } from "react";
import ContactsPrivate from "../contactprivate/contactprivate";
import { isAuthenticated } from "@/app/page";

export default function CreatePost({ onPostCreated }) {

    // let [privacy, setPrivacy] = useState("public")
    // let [title, setTitle] = useState("")
    // let [content, setContent] = useState("")
    // let [namePoster, setNamePoster] = useState("")
    // let [image, setImage] = useState(null)
    const [post,setPost] = useState({
        privacy:"public",
        title:"",
        content:"",
        namePoster:"",
        image: null
        
    })
    const fileInputRef = useRef(null)
    const host = process.env.NEXT_PUBLIC_HOST
    // const postData = {
    //     privacy: privacy,
    //     title: title,
    //     content: content
    // }


    const handleSubmit = async (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const postData = {
            privacy : formData.get('privacy'),
            title : formData.get('title'),
            content : formData.get('content'),
            image : formData.get('image')
            
        };
        // console.log("formData", postData);
        
        formData.append('postData', postData);  //JSON.stringify(postData)
        console.log("formData", formData);
        
        // if (image) {
        //     formData.append('avatar', image);
        // }

        try {
            const response = await fetch(`${host}/addPost`, {
                method: "POST",
                body: postData,
                credentials: "include"
            });

            // if (!postData.title.trim() || !postData.content.trim()) {
            //     return
            // }

            const newPost = await response.json();
            console.log("newPost", response);


            if (!response.ok) {
                isAuthenticated(response.status, newPost.error);
                throw new Error(newPost.error);
            } else {
                // Reset form
                // setPrivacy("public");
                // setTitle("");
                // setContent("");
                // setImage(null);

                setPost(p => p.title = "");
                setPost(p => p.content = "");
                setPost(p => p.image = null);
                setPost(p => p.privacy = "public");
            }

            // Reset file input
            if (fileInputRef.current) {
                fileInputRef.current.value = ""
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
                        {/* <Image
                            className={styles.image}
                            src="/images/post.png"
                            alt="post"
                            // width={500}
                            // height={500}
                            fill={true}
                        /> */}
                    </div>
                    <div className="nameProfile">
                        <h3>full name</h3>
                        <select /*onChange={(e) => { setPost(p=>p.privacy = e.target.value)  /*setPrivacy(e.target.value) }}*/ name="friends" id="friends" defaultValue={"public"}>
                            <option value={"public"}>Public</option>
                            <option value={"private"}>Private</option>
                            <option value={"almostPrivate"}>Almost Private </option>
                        </select>
                    </div>
                </div>
                {/* {privacy === "private" && <ContactsPrivate />} */}
                <div className="title">
                    <input /*onChange={(e) => { setPost(p) /*setTitle(e.target.value)* }} value={title}*/ type="text" name="title" placeholder="enter your title" />
                </div>
                <div className="content">
                    {/* <textarea name="content" placeholder="enter your content"></textarea> */}
                    <textarea
                        // onChange={(e) => { setContent(e.target.value) }}
                        // value={content}
                        placeholder="Write a content..."
                        name="content"
                        // value={comment}
                        // onChange={(e) => setComment(e.target.value)}
                        rows={4}
                        style={{ width: '90%', padding: '5px', borderRadius: '4px', resize: 'none', outline: 'none', border: 'none', marginLeft: '5%', backgroundColor: '#333' }}
                    />
                </div>
                <div className="uploadImage">
                    <input /*onChange={(e) => { setImage(e.target.files[0]) }}*/ id="uploadImage" className="hiddenInput" ref={fileInputRef} type="file" name="image" />
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
