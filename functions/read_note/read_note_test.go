package read_note

import (
	"context"
	"testing"

	"github.com/apotox/go-encrynote/functions/create_note"
	mocks "github.com/apotox/go-encrynote/mocks"
	pkg "github.com/apotox/go-encrynote/pkg"
	tools "github.com/apotox/go-encrynote/tools"
	"github.com/aws/aws-lambda-go/events"
	. "github.com/smartystreets/goconvey/convey"
)

func buildCreateEvent(body map[string]interface{}) events.APIGatewayProxyRequest {
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

func TestReadNote(t *testing.T) {

	t.Parallel()

	Convey("should call read_note function", t, func() {

		Convey("should return success code 1", func() {
			tools.CleanCollection(pkg.COL_NOTES)

			createResult, err := create_note.Handler(context.TODO(), buildCreateEvent(map[string]interface{}{
				"message":  "hello from test",
				"expireAt": "2022-07-31T10:16:55.732Z",
			}))

			So(err, ShouldBeNil)

			responseBody, parseError := pkg.ParseResponseBody(createResult)
			So(parseError, ShouldBeNil)

			So(responseBody, ShouldNotBeNil)

			id := responseBody["itemId"].(string) + responseBody["usedKey"].(string)

			queueDeleteNoteMocked := new(mocks.QueueDeleteNoteMock)
			queueDeleteNoteMocked.On("Init").Return(nil)
			queueDeleteNoteMocked.On("Publish", responseBody["itemId"].(string)).Return(nil)

			queueDeleteNoteMocked.Init()
			readResult, err := Handler(context.TODO(), buildReadEvent(id), pkg.Queues{
				QueueDeleteNote: queueDeleteNoteMocked,
			})
			So(err, ShouldBeNil)
			readResponseBody, parseError := pkg.ParseResponseBody(readResult)
			So(parseError, ShouldBeNil)
			So(readResponseBody, ShouldNotBeNil)
			So(readResult.StatusCode, ShouldEqual, 200)
		})
	})

}
