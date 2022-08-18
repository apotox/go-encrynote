package encryption

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncryption(t *testing.T) {

	t.Parallel()

	Convey("Encryption tests", t, func() {

		Convey("should fail if key has invalid length", func() {
			key := "test"
			_, err := Encrypt(key, "plaintext")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldContainSubstring, "crypto/aes: invalid key size 4")
		})

		Convey("should generate aes key", func() {
			key, err := GenerateAesKey()
			So(err, ShouldBeNil)
			So(len(key), ShouldEqual, 32)
		})

		Convey("should encrypt plaintext with nonce", func() {
			key, _ := GenerateAesKey()
			encrypted1, err := Encrypt(key, "plaintext")
			So(err, ShouldBeNil)
			So(len(encrypted1), ShouldBeGreaterThan, 0)

			encrypted2, err := Encrypt(key, "plaintext")
			So(err, ShouldBeNil)
			So(len(encrypted2), ShouldBeGreaterThan, 0)

			So(encrypted1, ShouldNotEqual, encrypted2)

		})

		Convey("should encrypt and decrypt plaintext successfully", func() {
			key, _ := GenerateAesKey()
			encrypted, err := Encrypt(key, "plaintext")
			So(err, ShouldBeNil)

			decrypted, err := Decrypt(key, encrypted)

			So(err, ShouldBeNil)
			So(decrypted, ShouldEqual, "plaintext")

		})

	})

}
