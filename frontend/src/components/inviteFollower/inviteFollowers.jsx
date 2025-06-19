

"use client";
import React, { useState, useEffect } from "react";
import "./invitefollowers.css"

const host = process.env.NEXT_PUBLIC_HOST;

export default function InviteUsers({ group_id }) {
    const [users, setUsers] = useState([]);
    const [error, setError] = useState("");
    const [invited, setInvited] = useState([]);

    async function GetUsers() {
        try {
            const responce = await fetch(
                `${host}/GetFolowingsUsers?groupId=${group_id}`,
                {
                    credentials: "include",
                    method: "GET",
                }
            );

            const data = await responce.json();
            console.log(data);

            if (!responce.ok) {
                setError(data.error);
                return;
            }

            setUsers(Array.isArray(data) ? data : []);
        } catch (err) {
            setError(err.message);
        }
    }

    async function InviteUser(user_id) {
        try {
            const responce = await fetch(`${host}/groupInvitarion`, {
                credentials: "include",
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    actor_id: parseInt(group_id),
                    target: user_id,
                }),
            });

            const data = await responce.json();
            if (!responce.ok) {
                setError(data.error);
                return;
            }

            setInvited((prev) => [...prev, user_id]);
        } catch (err) {
            setError("Error during invitation");
        }
    }

    useEffect(() => {
        GetUsers();
    }, []);

    return (
        <>
            <button className="soutitre0">invit users</button>

            <div className="userscontainer">
                {users.length > 0 ? (
                    users.map((user) => {
                        const isInvited = invited.includes(user.ID);
                        return (
                            <div key={user.ID} className="user-wrapper">
                                <div className="user">
                                    <img
                                        src={`http://${user.avatar}`}
                                        alt={`${user.firstname} ${user.lastname}`}
                                        className="avatar"
                                    />
                                    <p>
                                        {user.firstname} {user.lastname}
                                    </p>
                                </div>
                                <div className="invitation">
                                    <button
                                        className={isInvited ? "invited" : ""}
                                        onClick={() => InviteUser(user.ID)}
                                        disabled={isInvited}
                                    >
                                        {isInvited ? "Invited ✅" : "Invite user"}
                                    </button>
                                </div>
                            </div>
                        );
                    })
                ) : (
                    <p className="no-users">No users yet!</p>
                )}
            </div>

            <div className="error">{error}</div>
        </>
    );
}
