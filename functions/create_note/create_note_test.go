package create_note

import (
	"context"
	"testing"

	pkg "github.com/apotox/go-encrynote/pkg"
	tools "github.com/apotox/go-encrynote/tools"
	"github.com/aws/aws-lambda-go/events"
	. "github.com/smartystreets/goconvey/convey"
)

var defaultBody = map[string]interface{}{
	"message": "test",
}

func buildEvent(body map[string]interface{}) events.APIGatewayProxyRequest {
	return events.APIGatewayProxyRequest{
		Path:            "/note",
		HTTPMethod:      "POST",
		IsBase64Encoded: false,
		Body:            pkg.Marshal(body),
	}
}

func buildReadEvent(id string) events.APIGatewayProxyRequest {
	return events.APIGatewayProxyRequest{
		Path:            "/note",
		HTTPMethod:      "GET",
		IsBase64Encoded: false,
		QueryStringParameters: map[string]string{
			"id": id,
		},
	}
}

func TestCreateNote(t *testing.T) {

	t.Parallel()

	Convey("should call create_note function", t, func() {
		Convey("should return success code 1", func() {
			tools.CleanCollection(pkg.COL_NOTES)
			result, err := Handler(context.TODO(), buildEvent(defaultBody))

			So(err, ShouldBeNil)
			responseBody, parseError := pkg.ParseResponseBody(result)
			So(parseError, ShouldBeNil)
			So(responseBody, ShouldNotBeNil)
			So(result.StatusCode, ShouldEqual, 200)
			So(responseBody["itemId"], ShouldNotBeNil)
		})

		Convey("should throw if validation fail", func() {
			tools.CleanCollection(pkg.COL_NOTES)
			result, _ := Handler(context.TODO(), buildEvent(map[string]interface{}{
				"message": "",
			}))
			responseBody, parseError := pkg.ParseResponseBody(result)
			So(parseError, ShouldBeNil)
			So(responseBody, ShouldNotBeNil)
			So(result.StatusCode, ShouldEqual, 400)
			So(responseBody["error"], ShouldContainSubstring, "Message")
		})

		Convey("should throw if expireAt is invalid", func() {
			tools.CleanCollection(pkg.COL_NOTES)
			result, _ := Handler(context.TODO(), buildEvent(map[string]interface{}{
				"message":  "hello from test",
				"expireAt": "2022-0x7-31T10:16:55.732Z",
			}))

			response, parseError := pkg.ParseResponseBody(result)
			So(parseError, ShouldBeNil)
			So(response, ShouldNotBeNil)
			So(result.StatusCode, ShouldEqual, 500)
		})

		Convey("should save the encrypted message", func() {
			tools.CleanCollection(pkg.COL_NOTES)
			message := "hello from test"
			createResult, err := Handler(context.TODO(), buildEvent(map[string]interface{}{
				"message":  message,
				"expireAt": "2022-07-31T10:16:55.732Z",
			}))
			So(err, ShouldBeNil)

			createResponseBody, parseError := pkg.ParseResponseBody(createResult)
			So(parseError, ShouldBeNil)

			So(createResponseBody, ShouldNotBeNil)
			So(createResult.StatusCode, ShouldEqual, 200)

			noteId := createResponseBody["itemId"].(string)
			note, err := tools.Read[pkg.Note](pkg.COL_NOTES, noteId)

			So(err, ShouldBeNil)
			So(note, ShouldNotBeNil)
			So(note.Message, ShouldNotBeEmpty)
			So(note.Message, ShouldNotEqual, message)
		})

	})

}
