/* // "client use";
import Link from "next/link";
import "./groupcard.css";
import { useEffect, useState } from "react";
export default function GroupCard({ groups }) {
  console.log(groups);

  const [currentGroup, setCurrentGroup] = useState(null);
  const host = process.env.NEXT_PUBLIC_HOST;
  async function joinReq() {
    if (!currentGroup) return;
    const resp = await fetch(`${host}/JouindGroupe`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: { groupe_id: currentGroup },
    });
    const data = resp.json();
    console.log(data);
  }

  useEffect(() => {
    joinReq();
    alert(`${host}/JouindGroupe`);
  }, [currentGroup]);
  return (
    <>
      {
        <div className="groupscard">
          {groups.map((groupe, i) => (
            <div className="card" key={i}>
              <h2 className="groupTitle">{groupe.title}</h2>
              <label className="grouplabel">group description:</label>
              <p className="groupDescription">{groupe.description}</p>
              {groupe.status === "member" ? (
                <>
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
                    enter to group
                  </Link>
                </>
              ) : groupe.status === "requested" ? (
                <>
                  <button className="requestbtn">requested</button>
                </>
              ) : (
                <>
                  <button
                    className="joinreqbtn"
                    onClick={() => {
                      console.log(groupe);

                      alert("the function get fired ", groupe.Id);
                      setCurrentGroup(groupe.id);
                    }}
                  >
                    join request
                  </button>
                </>
              )}
            </div>
          ))}
        </div>
      }
    </>
  );
}
 */

// "client use";
import Link from "next/link";
import "./groupcard.css";
import { useEffect, useState } from "react";

export default function GroupCard({ groups }) {
  const [currentGroup, setCurrentGroup] = useState(null);
  const [groupsStat, setGroupsStat] = useState(groups);
  const host = process.env.NEXT_PUBLIC_HOST;

  async function joinReq() {
    if (!currentGroup) return;

    try {
      const resp = await fetch(`${host}/JouindGroupe`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ groupe_id: currentGroup.Id }),
      });

      if (!resp.ok) {
        throw new Error("Failed to send join request");
      }

      const data = await resp.json();
      console.log(data);
      if (data.prossotion === "succeeded") {
        setGroupsStat(groups);
      }
    } catch (error) {
      console.error("Error:", error);
    }
  }

  useEffect(() => {
    if (currentGroup) {
      joinReq();
    }
  }, [currentGroup]);

  return (
    <div className="groupscard">
      {groups.map((groupe, i) => (
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
            <button
              className="joinreqbtn"
              onClick={() => {
                console.log(groupe);
                setCurrentGroup(groupe);
              }}
            >
              Join request
            </button>
          )}
        </div>
      ))}
    </div>
  );
}
