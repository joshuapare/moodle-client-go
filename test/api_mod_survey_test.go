/*
Moodle Webservice API

Testing ModSurveyAPIService

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

func Test_moodleclient_ModSurveyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ModSurveyAPIService ModSurveyGetQuestions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ModSurveyAPI.ModSurveyGetQuestions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModSurveyAPIService ModSurveyGetSurveysByCourses", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ModSurveyAPI.ModSurveyGetSurveysByCourses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModSurveyAPIService ModSurveySubmitAnswers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ModSurveyAPI.ModSurveySubmitAnswers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ModSurveyAPIService ModSurveyViewSurvey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ModSurveyAPI.ModSurveyViewSurvey(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
