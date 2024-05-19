/*
Moodle Webservice API

Testing QbankTagquestionAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package moodleclient

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/joshuapare/moodle-client-go"
)

func Test_moodleclient_QbankTagquestionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test QbankTagquestionAPIService QbankTagquestionSubmitTagsForm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.QbankTagquestionAPI.QbankTagquestionSubmitTagsForm(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}