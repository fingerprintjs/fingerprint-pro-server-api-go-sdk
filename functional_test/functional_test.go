package functional_test

import (
	"context"
	"github.com/fingerprintjs/fingerprint-pro-server-api-go-sdk/v7/sdk"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestApiFunctional(t *testing.T) {
	// Load environment variables
	godotenv.Load()

	apiKey := os.Getenv("FINGERPRINT_API_KEY")
	assert.NotEmpty(t, apiKey)
	auth := context.WithValue(context.Background(), sdk.ContextAPIKey, sdk.APIKey{
		Key: apiKey,
	})
	cfg := sdk.NewConfiguration()
	client := sdk.NewAPIClient(cfg)

	end := time.Now().UnixMilli()
	start := time.Now().AddDate(0, 0, -90).UnixMilli()
	opts := sdk.FingerprintApiSearchEventsOpts{
		Start: &start,
		End:   &end,
	}
	events, _, err := client.FingerprintApi.SearchEvents(auth, 2, &opts)
	assert.NoError(t, err)
	assert.NotNil(t, events.Events)
	testEvent := events.Events[0]
	requestId := testEvent.Products.Identification.Data.RequestId
	visitorId := testEvent.Products.Identification.Data.VisitorId

	t.Run("GetEvent", func(t *testing.T) {
		t.Run("with valid event", func(t *testing.T) {
			event, _, err := client.FingerprintApi.GetEvent(auth, requestId)

			assert.NoError(t, err)
			assert.NotNil(t, event.Products)
			assert.IsType(t, &sdk.Products{}, event.Products)
			assert.IsType(t, &sdk.ProductIdentification{}, event.Products.Identification)
			assert.Equal(t, requestId, event.Products.Identification.Data.RequestId)
			assert.Equal(t, visitorId, event.Products.Identification.Data.VisitorId)
		})

		t.Run("error 404", func(t *testing.T) {
			event, _, err := client.FingerprintApi.GetEvent(auth, "1662542583652.pLBzes")

			assert.Error(t, err)
			assert.Nil(t, event.Products)
			assert.Equal(t, sdk.REQUEST_NOT_FOUND, err.Code())
			assert.IsType(t, &sdk.ErrorResponse{}, err.Model())
		})
	})

	t.Run("GetVisits", func(t *testing.T) {
		t.Run("without request id", func(t *testing.T) {
			opts := sdk.FingerprintApiGetVisitsOpts{}
			visit, _, err := client.FingerprintApi.GetVisits(auth, visitorId, &opts)

			assert.NoError(t, err)
			assert.NotNil(t, visit)
			assert.Equal(t, visitorId, visit.VisitorId)
			assert.GreaterOrEqual(t, len(visit.Visits), 1)
		})

		t.Run("with request id", func(t *testing.T) {
			opts := sdk.FingerprintApiGetVisitsOpts{
				RequestId: requestId,
			}
			visit, _, err := client.FingerprintApi.GetVisits(auth, visitorId, &opts)

			assert.NoError(t, err)
			assert.NotNil(t, visit)
			assert.Equal(t, visitorId, visit.VisitorId)
			assert.Equal(t, requestId, visit.Visits[0].RequestId)
			assert.Len(t, visit.Visits, 1)
		})
	})

	t.Run("SearchEvents", func(t *testing.T) {
		t.Run("simple search", func(t *testing.T) {
			end := time.Now().UnixMilli()
			start := time.Now().AddDate(0, 0, -365).UnixMilli()
			opts := sdk.FingerprintApiSearchEventsOpts{
				Start: &start,
				End:   &end,
			}
			events, _, err := client.FingerprintApi.SearchEvents(auth, 2, &opts)
			assert.NoError(t, err)
			assert.NotNil(t, events.Events)
			assert.Len(t, events.Events, 2)
		})

		t.Run("with pagination", func(t *testing.T) {
			end := time.Now().UnixMilli()
			start := time.Now().AddDate(0, 0, -365).UnixMilli()
			events, _, err := client.FingerprintApi.SearchEvents(auth, 2, &sdk.FingerprintApiSearchEventsOpts{
				Start: &start,
				End:   &end,
			})
			assert.NoError(t, err)
			assert.NotNil(t, events.Events)
			assert.NotEmpty(t, events.PaginationKey)
			assert.Len(t, events.Events, 2)

			nextEvents, _, err := client.FingerprintApi.SearchEvents(auth, 2, &sdk.FingerprintApiSearchEventsOpts{
				Start:         &start,
				End:           &end,
				PaginationKey: &events.PaginationKey,
			})
			assert.NoError(t, err)
			assert.NotNil(t, nextEvents)
			assert.Len(t, nextEvents.Events, 2)
			assert.NotEqual(t, events.Events[0].Products.Identification.Data.RequestId, nextEvents.Events[0].Products.Identification.Data.RequestId)
			assert.NotEqual(t, events.Events[1].Products.Identification.Data.RequestId, nextEvents.Events[1].Products.Identification.Data.RequestId)
		})

		t.Run("with old events", func(t *testing.T) {
			end := time.Now().UnixMilli()
			start := time.Now().AddDate(0, 0, -365).UnixMilli()
			reverse := true
			events, _, err := client.FingerprintApi.SearchEvents(auth, 2, &sdk.FingerprintApiSearchEventsOpts{
				Start:   &start,
				End:     &end,
				Reverse: &reverse,
			})
			assert.NoError(t, err)
			assert.NotNil(t, events.Events)
			assert.Len(t, events.Events, 2)

			oldEvent := events.Events[0]

			// Try to get old events to check if they still could be deserialized
			oldVisit, _, err := client.FingerprintApi.GetVisits(auth, oldEvent.Products.Identification.Data.RequestId, nil)
			assert.NoError(t, err)
			assert.NotNil(t, oldEvent.Products)
			assert.Equal(t, oldEvent.Products.Identification.Data.RequestId, oldVisit.VisitorId)

			oldGetEvent, _, err := client.FingerprintApi.GetEvent(auth, oldEvent.Products.Identification.Data.RequestId)
			assert.NoError(t, err)
			assert.NotNil(t, oldGetEvent.Products)
			assert.Equal(t, oldEvent.Products.Identification.Data.RequestId, oldGetEvent.Products.Identification.Data.RequestId)
		})
	})
}
