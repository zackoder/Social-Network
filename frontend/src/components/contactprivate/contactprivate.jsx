"use client";
import React, { useContext, useState, useEffect } from "react";
import "./contactprivate.modules.css";








  const host = process.env.NEXT_PUBLIC_HOST;







function displayChatbox() {
  const button = document.querySelector(".soutitre0");
  const container = document.querySelector(".userscontainer");
  // const formSubmit = document.querySelector(".submitForm");

  button?.addEventListener("click", () => {
    if (container?.classList.contains("showw")) {
      container.classList.remove("showw");
      container.classList.add("hide");
      // formSubmit?.classList.add("hide");

      // After animation ends, hide the element
      button.addEventListener(
        "click",
        () => {
          if (container?.classList.contains("hide")) {
            container.style.display = "none";
            // formSubmit.style.display = "none";
          }
        },
        { once: true }
      );
    } else {
      container.classList.remove("hide");
      container.classList.add("showw");
      // formSubmit?.classList.add("showw");
      container.style.display = "block";
    }
  });
}


export default function InviteUsers(group_id) {
  let [users, setusers] = useState([]);
  let [error,Seterror]=useState("");
   async function GetUsers() {
  let responce = await fetch(`${host}/GetFolowingsUsers`,{
    credentials:"include",
    body:JSON.stringify(group_id)
  });

  const data = await responce.json();
    if (!responce.ok){
      Seterror(data.message)
      return
}
  setusers(data)



}
  return <>
    <button className="soutitre0" onClick={displayChatbox}>
      invit users
    </button>

 <div className="userscontainer">
  {users.length > 0 ? (
    users.map((user) => (
      <div key={user.id} className="user-wrapper">
        <div className="user">
          <p>{user.firstname} {user.lastname}</p>
        </div>
        <div className="invitation">
          <button>Invite user</button>
        </div>
      </div>
    ))
  ) : (
    <div>
      <p>No users yet</p>
    </div>
  )}
</div>
<div className="error">{error}</div>

  </>
}
