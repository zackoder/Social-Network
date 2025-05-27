export default function GroupPage({ params }) {
  // params est injecté automatiquement en mode serveur dans /app
  const groupId = params.id;
  console.log(params);


  return (
    <div>
      <h1>Page du groupe</h1>

      <p>ID du groupe : {groupId}</p>
    </div>
  );
}