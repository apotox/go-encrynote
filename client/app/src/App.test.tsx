import React from "react";
import { render, screen } from "@testing-library/react";
import App from "./App";
import { BrowserRouter } from "react-router-dom";
import { RecoilRoot } from "recoil";

test("renders Encrynote title", () => {
	render(
		<BrowserRouter>
			<RecoilRoot>
				<App />
			</RecoilRoot>
		</BrowserRouter>
	);
	const linkElement = screen.getByText(/Encrynote/i);
	expect(linkElement).toBeInTheDocument();
});
