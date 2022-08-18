import { Snackbar } from "@mui/material";
import { useRecoilValue } from "recoil";
import { toastAtom } from "./atom";
export { useToast } from "./useToast";


export const Toast = () => {

  const state = useRecoilValue(toastAtom);

  const handleClose = () => {};

  return (
    <Snackbar
      open={state.open}
      autoHideDuration={3000}
      onClose={handleClose}
      message={state.message}
    />
  );
};
