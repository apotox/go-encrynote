import React from "react";
import { Typography } from "@mui/material";
import { Suspense, lazy } from "react";
import Layout from "../components/Layout";

const LazyEncrytion = lazy(() => import("../components/Encryption"));

export default function Home() {
	return (
		<Layout>
			<Typography
				sx={{ fontSize: 14, marginBottom: 4 }}
				color="text.secondary"
				gutterBottom
			>
        encrypt your notes
			</Typography>

			<Suspense fallback={<p>initialling..</p>}>
				<LazyEncrytion />
			</Suspense>
		</Layout>
	);
}
