package pkg

import (
	"fmt"
	"strings"
	"time"

	aes "github.com/apotox/go-encrynote/encryption"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UnmarshalJSON parse isoformat date
func (t *Datetime) UnmarshalJSON(input []byte) error {

	if strings.Contains(string(input), "null") {
		return nil
	}

	strInput := strings.Trim(string(input), `"`)
	newTime, err := time.Parse(time.RFC3339, strInput)
	if err != nil {
		return err
	}

	t.DateTime = primitive.NewDateTimeFromTime(newTime)

	return nil
}

// GetCollection returns the collection name for the given entity
func (n *Note) GetCollectionName() string {
	return string(COL_NOTES)
}

// EncryptMessage encrypts the message with random key and return the key
func (n *Note) EncryptMessage() string {
	key, _ := aes.GenerateAesKey()
	encryptedMessage, _ := aes.Encrypt(key, n.Message)
	n.Message = encryptedMessage
	return key
}

func (n *Note) DecryptMessage(key string) string {
	DecryptedMessage, _ := aes.Decrypt(key, n.Message)
	return DecryptedMessage
}

// generate and set note URL
func (n *Note) GenerateUrl(usedKey string) string {

	n.URL = fmt.Sprintf("note/%s%s", n.Id, usedKey)

	return n.URL
}
