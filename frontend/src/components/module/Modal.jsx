import React from "react";
import "./Modal.css"; // Assure-toi que le chemin est correct

const Modal = ({ isOpen, onClose, children }) => {
  if (!isOpen) return null;

  return (
    <div className="modalOverlay">
      <div className="modalContent">
        <button className="closeButton" onClick={onClose}>X</button>
        {children}
      </div>
    </div>
  );
};

export default Modal;
