/*
Stacks Blockchain API

Testing NamesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package stacks_blockchain_api_client

import (
	"context"
	"testing"

	"github.com/icon-project/stacks-go-sdk/pkg/stacks_blockchain_api_client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_stacks_blockchain_api_client_NamesAPIService(t *testing.T) {

	configuration := stacks_blockchain_api_client.NewConfiguration()
	apiClient := stacks_blockchain_api_client.NewAPIClient(configuration)

	t.Run("Test NamesAPIService FetchSubdomainsListForName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.NamesAPI.FetchSubdomainsListForName(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService FetchZoneFile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.NamesAPI.FetchZoneFile(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetAllNames", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.NamesAPI.GetAllNames(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetAllNamespaces", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.NamesAPI.GetAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetHistoricalZoneFile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string
		var zoneFileHash string

		resp, httpRes, err := apiClient.NamesAPI.GetHistoricalZoneFile(context.Background(), name, zoneFileHash).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNameInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.NamesAPI.GetNameInfo(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNamePrice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.NamesAPI.GetNamePrice(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNamesOwnedByAddress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var blockchain string
		var address string

		resp, httpRes, err := apiClient.NamesAPI.GetNamesOwnedByAddress(context.Background(), blockchain, address).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNamespaceNames", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tld string

		resp, httpRes, err := apiClient.NamesAPI.GetNamespaceNames(context.Background(), tld).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NamesAPIService GetNamespacePrice", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var tld string

		resp, httpRes, err := apiClient.NamesAPI.GetNamespacePrice(context.Background(), tld).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
