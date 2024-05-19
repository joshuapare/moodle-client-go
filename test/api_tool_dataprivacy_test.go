/*
Moodle Webservice API

Testing ToolDataprivacyAPIService

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

func Test_moodleclient_ToolDataprivacyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyApproveDataRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyApproveDataRequest(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyBulkApproveDataRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyBulkApproveDataRequests(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyBulkDenyDataRequests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyBulkDenyDataRequests(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyCancelDataRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyCancelDataRequest(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyConfirmContextsForDeletion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyConfirmContextsForDeletion(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyContactDpo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyContactDpo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyCreateCategoryForm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyCreateCategoryForm(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyCreatePurposeForm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyCreatePurposeForm(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyDeleteCategory", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyDeleteCategory(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyDeletePurpose", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyDeletePurpose(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyDenyDataRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyDenyDataRequest(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyGetActivityOptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyGetActivityOptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyGetCategoryOptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyGetCategoryOptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyGetDataRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyGetDataRequest(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyGetPurposeOptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyGetPurposeOptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyGetUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyGetUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyMarkComplete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyMarkComplete(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacySetContextDefaults", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacySetContextDefaults(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacySetContextForm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacySetContextForm(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacySetContextlevelForm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacySetContextlevelForm(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacySubmitSelectedCoursesForm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacySubmitSelectedCoursesForm(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ToolDataprivacyAPIService ToolDataprivacyTreeExtraBranches", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ToolDataprivacyAPI.ToolDataprivacyTreeExtraBranches(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
