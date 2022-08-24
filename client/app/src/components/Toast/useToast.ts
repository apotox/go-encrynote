import { useRecoilState } from "recoil";
import { toastAtom } from "./atom";



export const useToast = () => {
	const [toast, setToast] = useRecoilState(toastAtom);
    
	const open = (message: string, type: "success" | "error" | "info") => {
		setToast({
			message,
			type,
			open: true,
		});
		setTimeout(() => {
			close();
		}, 3000);
	}
    
	const close = () => {
		setToast({
			...toast,
			open: false,
		});
	}
    
	return {
		toast,
		open,
		close,
	};
}