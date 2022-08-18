package delete_note

import (
	"context"
	"testing"

	"github.com/apotox/go-encrynote/functions/create_note"
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

func buildDeleteEvent(id string) events.SQSEvent {
	return events.SQSEvent{
		Records: []events.SQSMessage{
			{
				Body: id,
			},
		},
	}
}

func TestDeleteNote(t *testing.T) {

	t.Parallel()

	Convey("should call delete_note function", t, func() {

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

			id := responseBody["itemId"].(string)

			err = Handler(context.TODO(), buildDeleteEvent(id))
			So(err, ShouldBeNil)

		})
	})

}
