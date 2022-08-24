import { atom } from "recoil";

export interface ToastState {
  message: string;
  type: "success" | "error" | "info";
  open: boolean;
}

export const toastAtom = atom({
	key: "toast",
	default: {
		message: "",
		type: "info",
		open: false,
	},
});
