import styles from "./groups.module.css";


import { useState } from 'react';

export default function Groupes() {
  // État initial pour tous les groupes, ceux que l'utilisateur a créés, et ceux dont il est membre
  const [groupes, setGroupes] = useState([
    { id: 1, nom: 'Développeurs Web', description: 'Un groupe de passionnés de développement web.', createur: true, membre: true },
    { id: 2, nom: 'Designers', description: 'Les experts en design graphique et UX/UI.', createur: false, membre: true },
    { id: 3, nom: 'Photographes', description: 'Un groupe pour les photographes amateurs et professionnels.', createur: true, membre: false },
    { id: 4, nom: 'Musiciens', description: 'Groupe pour les passionnés de musique.', createur: false, membre: false },
  ]);

  // Notifications (simulation)
  const [notifications, setNotifications] = useState([
    { id: 1, groupeId: 1, message: 'Nouvelle mise à jour de la documentation.' },
    { id: 2, groupeId: 2, message: 'Réunion prévue demain à 18h.' },
  ]);

  // Ajouter un nouveau groupe
  const [nom, setNom] = useState('');
  const [description, setDescription] = useState('');

  const handleAjouterGroupe = () => {
    if (nom && description) {
      setGroupes([...groupes, { id: groupes.length + 1, nom, description, createur: true, membre: true }]);
      setNom('');
      setDescription('');
    }
  };

  return (
    <div>
      <h1>Gestion des Groupes</h1>

      {/* Section 1: Tous les groupes */}
      <section>
        <h2>Tous les Groupes</h2>
        <ul>
          {groupes.map(groupe => (
            <li key={groupe.id}>
              <h3>{groupe.nom}</h3>
              <p>{groupe.description}</p>
            </li>
          ))}
        </ul>
      </section>

      {/* Section 2: Groupes créés par l'utilisateur */}
      <section>
        <h2>Groupes que j'ai créés</h2>
        <ul>
          {groupes.filter(groupe => groupe.createur).map(groupe => (
            <li key={groupe.id}>
              <h3>{groupe.nom}</h3>
              <p>{groupe.description}</p>
            </li>
          ))}
        </ul>
      </section>

      {/* Section 3: Groupes dont l'utilisateur est membre */}
      <section>
        <h2>Groupes dont je suis membre</h2>
        <ul>
          {groupes.filter(groupe => groupe.membre).map(groupe => (
            <li key={groupe.id}>
              <h3>{groupe.nom}</h3>
              <p>{groupe.description}</p>
            </li>
          ))}
        </ul>
      </section>

      {/* Section 4: Notifications des groupes */}
      <section>
        <h2>Notifications des Groupes</h2>
        <ul>
          {notifications.map(notification => (
            <li key={notification.id}>
              <strong>Groupe {groupes.find(g => g.id === notification.groupeId)?.nom}</strong>: {notification.message}
            </li>
          ))}
        </ul>
      </section>

      {/* Formulaire pour ajouter un groupe */}
      <section>
        <h2>Ajouter un nouveau groupe</h2>
        <input
          type="text"
          placeholder="Nom du groupe"
          value={nom}
          onChange={(e) => setNom(e.target.value)}
        />
        <textarea
          placeholder="Description du groupe"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <button onClick={handleAjouterGroupe}>Ajouter</button>
      </section>


    </div>
  );
}
