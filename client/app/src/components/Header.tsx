import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";

export default function Header() {
	return (
		<AppBar position="static">
			<Toolbar>
				<Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
					<a
						style={{
							color: "white",
							textDecoration: "none",
						}}
						href="/"
					>
            Encrynote
					</a>
				</Typography>
				<Button color="inherit" href="https://github.com/apotox/go-encrynote">
          Github
				</Button>
			</Toolbar>
		</AppBar>
	);
}
