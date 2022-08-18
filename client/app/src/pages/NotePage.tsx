import React, { lazy } from "react";
import { Typography } from "@mui/material";
import { useParams } from "react-router-dom";
import Layout from "../components/Layout";
const Decryption = lazy(() => import("../components/Decryption"));

export default function NotePage() {
  const { id } = useParams<string>();

  return (
    <Layout>
      <h2>Note</h2>
      <Typography
        sx={{ fontSize: 14, marginBottom: 4 }}
        color="text.secondary"
        gutterBottom
      >
        decrypt your notes {id}
      </Typography>

      <Decryption id={id || ""} />
    </Layout>
  );
}
