import React from "react";
import getConfig from "../config";

export default function Footer() {
  return (
    <footer
      style={{
        position: "fixed",
        bottom: 0,
        left: 0,
        right: 0,
        backgroundColor: "#f5f5f5",
        padding: "1rem",
        textAlign: "center",
      }}
    >
      <span>Commit sha: {getConfig().REACT_APP_APP_SHORTSHA}</span>
    </footer>
  );
}
