"use client";
import Link from "next/link";
import "./groupcard.css";
import { use, useEffect, useState } from "react";

export default function GroupCard({ groups }) {
  const [groupsData, setGroupStatuses] = useState([]);

  useEffect(() => {
    setGroupStatuses(groups);
  }, [groups]);

  const host = process.env.NEXT_PUBLIC_HOST;

  async function joinReq(groupe) {
    try {
      const resp = await fetch(`${host}/JouindGroupe`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ groupe_id: groupe.Id }),
      });

      if (!resp.ok) {
        throw new Error("Failed to send join request");
      }

      const data = await resp.json();
      console.log(data);

      if (data.prossotion === "succeeded") {
        setGroupStatuses(
          groupsData.map((g) =>
            g.Id === groupe.Id ? { ...g, status: "requested" } : g
          )
        );
      }
    } catch (error) {
      console.error("Join request error:", error);
    }
  }

  return (
    <div className="groupscard">
      {groupsData.map((groupe, i) => {
        return (
          <div className="card" key={i}>
            <h2 className="groupTitle">{groupe.title}</h2>
            <label className="grouplabel">Group description:</label>
            <p className="groupDescription">{groupe.description}</p>

            {groupe.status === "member" ? (
              <Link
                className="enterGroup"
                href={{
                  pathname: `/groups/${groupe.Id}`,
                  query: {
                    Id: groupe.Id,
                    title: groupe.title,
                    description: groupe.description,
                  },
                }}
              >
                Enter the group
              </Link>
            ) : groupe.status === "requested" ? (
              <button className="requestbtn" disabled>
                Requested
              </button>
            ) : (
              <button className="joinreqbtn" onClick={() => joinReq(groupe)}>
                Join request
              </button>
            )}
          </div>
        );
      })}
    </div>
  );
}
