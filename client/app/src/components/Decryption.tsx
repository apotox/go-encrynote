import React, { useEffect, useState } from "react";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Button from "@mui/material/Button";
import { TextField } from "@mui/material";
import Client from "../axios-client";
import { Note } from "../common/types";

function Decryption({ id }: { id: string }) {
	const [message, setMessage] = useState<string>("");
	const [isLoading, setLoading] = useState(false);
	const [errorMessage, setErrorMessage] = useState("");

	const getMessage = async () => {
		setLoading(true);
		Client.get<Note>(`note/?id=${id}`)
			.then((response) => {
				setMessage(response.data.message);
			})
			.catch((error) => {
				console.error(error);
				setErrorMessage("note deleted or not found");
			})
			.finally(() => {
				setLoading(false);
			});
	};

	useEffect(() => {
		getMessage();
	}, []);

	if (isLoading) {
		return <div>Loading...</div>;
	}

	return (
		<Card sx={{ minWidth: 275, margin: "0 auto" }} elevation={1}>
			<CardContent>
				<TextField
					id="outlined-multiline-flexible"
					label="your note"
					multiline
					style={{ width: "100%" }}
					maxRows={10}
					value={message}
					helperText={errorMessage ?? ""}
				/>
			</CardContent>
			<CardActions
				style={{
					justifyContent: "flex-end",
				}}
			>
				<Button
					onClick={() => {
						navigator.clipboard.writeText(message);
					}}
					size="small"
				>
          Copy
				</Button>
			</CardActions>
		</Card>
	);
}

export default Decryption;
