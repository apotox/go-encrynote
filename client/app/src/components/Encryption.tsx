import * as React from "react";
import { useState } from "react";
import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Button from "@mui/material/Button";
import { Checkbox, FormControlLabel, TextField } from "@mui/material";
import PasswordInput from "./PasswordInput";
import axiosClient from "../axios-client";
import { useToast } from "./Toast";
import getConfig from "../config";

interface EncryptResult {
  usedKey: string;
  itemId: string;
}

export default function Encryption() {
  const [message, setMessage] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [isPasswordUsed, setIsPasswordUsed] = useState<boolean>(false);
  const [deleteOnRead, setDeleteOnRead] = useState<boolean>(false);
  const [isLoading, setLoading] = useState<boolean>(false);

  const [encryptResult, setEncryptResult] = useState<EncryptResult | null>(
    null
  );

  const toast = useToast();

  const handleEncrypt = async () => {
    setLoading(true);
    axiosClient
      .post<EncryptResult>("/note", {
        message,
      })
      .then((response) => {
        setMessage("");
        setPassword("");
        toast.open("Note encrypted", "success");
        setEncryptResult(response.data);
      })
      .catch((error) => {
        toast.open(error.message || "something wrong!", "error");
      })
      .finally(() => {
        setLoading(false);
      });
  };

  if (isLoading) {
    return <div>Loading...</div>;
  }

  const generateLink = () => {
    if (!encryptResult) {
      return "";
    }
    const id = `${encryptResult.itemId}${encryptResult.usedKey}`;
    return getConfig().REACT_APP_URL + "/#/note/" + id;
  };

  return (
    <>
      <Card sx={{ minWidth: 275, margin: "0 auto" }} elevation={1}>
        <CardContent>
          <TextField
            id="outlined-multiline-flexible"
            label="your note"
            multiline
            style={{ width: "100%" }}
            maxRows={10}
            inputProps={{
              maxLength: 140,
            }}
            helperText={`${
              message.length > 0 ? message.length + "/" : ""
            }140 characters max`}
            value={message}
            onChange={(e) => setMessage(e.target.value)}
          />

          <Box
            sx={{
              marginTop: "1rem",
            }}
          >
            <FormControlLabel
              control={
                <Checkbox
                  checked={isPasswordUsed}
                  disabled={true}
                  onChange={(e) => setIsPasswordUsed(!isPasswordUsed)}
                />
              }
              label={"protect this note with a password?"}
            />
            {isPasswordUsed && (
              <PasswordInput password={password} setPassword={setPassword} />
            )}
          </Box>

          <Box
            sx={{
              marginTop: "0.4rem",
            }}
          >
            <FormControlLabel
              control={
                <Checkbox
                  disabled={true}
                  checked={deleteOnRead}
                  onChange={(e) => setDeleteOnRead(!deleteOnRead)}
                />
              }
              label={"delete note after read?"}
            />
          </Box>
        </CardContent>
        <CardActions
          style={{
            justifyContent: "flex-end",
          }}
        >
          <Button
            size="small"
            onClick={handleEncrypt}
            disabled={isLoading || !message}
          >
            Encrypt
          </Button>
        </CardActions>
      </Card>

      {encryptResult && (
        <Card style={styles.card}>
          <CardContent
            style={{
              display: "flex",
              flexDirection: "column",
            }}
          >
            <TextField
              fullWidth
              id="outlined-multiline-flexible"
              label="your secret key"
              value={encryptResult.usedKey}
              style={{
                margin: "0.5rem",
              }}
              onClick={(e) => {
                (e.target as HTMLInputElement).select();
              }}
            />
            <TextField
              fullWidth
              id="outlined-multiline-flexible"
              label="your message id"
              value={encryptResult.itemId}
              style={{
                margin: "0.5rem",
              }}
              onClick={(e) => {
                (e.target as HTMLInputElement).select();
              }}
            />

            <TextField
              fullWidth
              id="outlined-multiline-flexible"
              label="share this link with your friend"
              value={generateLink()}
              style={{
                margin: "0.5rem",
              }}
              onClick={(e) => {
                (e.target as HTMLInputElement).select();
              }}
            />
          </CardContent>
        </Card>
      )}
    </>
  );
}

const styles = {
  card: {
    marginTop: "1rem",
  },
};
