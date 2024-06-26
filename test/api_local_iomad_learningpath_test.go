/*
Moodle Webservice API

Testing LocalIomadLearningpathAPIService

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

func Test_moodleclient_LocalIomadLearningpathAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LocalIomadLearningpathAPIService LocalIomadLearningpathActivate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocalIomadLearningpathAPI.LocalIomadLearningpathActivate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocalIomadLearningpathAPIService LocalIomadLearningpathAddcourses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocalIomadLearningpathAPI.LocalIomadLearningpathAddcourses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocalIomadLearningpathAPIService LocalIomadLearningpathAddusers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocalIomadLearningpathAPI.LocalIomadLearningpathAddusers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocalIomadLearningpathAPIService LocalIomadLearningpathCopypath", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocalIomadLearningpathAPI.LocalIomadLearningpathCopypath(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocalIomadLearningpathAPIService LocalIomadLearningpathDeletepath", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocalIomadLearningpathAPI.LocalIomadLearningpathDeletepath(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocalIomadLearningpathAPIService LocalIomadLearningpathGetcourses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocalIomadLearningpathAPI.LocalIomadLearningpathGetcourses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocalIomadLearningpathAPIService LocalIomadLearningpathGetprospectivecourses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocalIomadLearningpathAPI.LocalIomadLearningpathGetprospectivecourses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocalIomadLearningpathAPIService LocalIomadLearningpathGetprospectiveusers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocalIomadLearningpathAPI.LocalIomadLearningpathGetprospectiveusers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocalIomadLearningpathAPIService LocalIomadLearningpathGetusers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocalIomadLearningpathAPI.LocalIomadLearningpathGetusers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocalIomadLearningpathAPIService LocalIomadLearningpathOrdercourses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocalIomadLearningpathAPI.LocalIomadLearningpathOrdercourses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocalIomadLearningpathAPIService LocalIomadLearningpathRemovecourses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocalIomadLearningpathAPI.LocalIomadLearningpathRemovecourses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocalIomadLearningpathAPIService LocalIomadLearningpathRemoveusers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocalIomadLearningpathAPI.LocalIomadLearningpathRemoveusers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
