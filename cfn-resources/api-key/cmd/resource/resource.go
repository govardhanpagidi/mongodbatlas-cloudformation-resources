// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	atlasSDK "go.mongodb.org/atlas-sdk/v20230201002/admin"
)

var CreateRequiredFields = []string{constants.OrgID, constants.Description}
var UpdateRequiredFields = []string{constants.OrgID, constants.ID}
var ReadRequiredFields = []string{constants.OrgID, constants.ID}
var DeleteRequiredFields = []string{constants.OrgID, constants.ID}
var ListRequiredFields = []string{constants.OrgID}

const (
	CREATE = "CREATE"
	READ   = "READ"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
	LIST   = "LIST"
)

func setup() {
	util.SetupLogger("mongodb-atlas-api-key")
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(CreateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	//Set the roles from model
	apiKeyInput := atlasSDK.CreateAtlasOrganizationApiKey{
		Desc:  currentModel.Description,
		Roles: currentModel.Roles,
	}
	apiKeyRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.CreateApiKey(
		context.Background(),
		*currentModel.OrgId,
		&apiKeyInput,
	)
	apiKeyUserDetails, response, err := apiKeyRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		return handleError(response, CREATE, err)
	}

	//Read response
	model := currentModel.readAPIKeyDetails(apiKeyUserDetails)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Completed",
		ResourceModel:   model}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ReadRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	apiKeyRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.GetApiKey(
		context.Background(),
		*currentModel.OrgId,
		*currentModel.Id,
	)
	apiKeyUserDetails, response, err := apiKeyRequest.Execute()

	defer closeResponse(response)
	if err != nil {
		return handleError(response, READ, err)
	}

	model := currentModel.readAPIKeyDetails(apiKeyUserDetails)
	_, _ = logger.Debugf("Read Response: %+v", currentModel)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
		ResourceModel:   model}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(UpdateRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	//Set the roles from model
	apiKeyInput := atlasSDK.CreateAtlasOrganizationApiKey{
		Desc:  currentModel.Description,
		Roles: currentModel.Roles,
	}
	updateRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.UpdateApiKey(
		context.Background(),
		*currentModel.OrgId,
		*currentModel.Id,
		&apiKeyInput,
	)
	apiKeyUserDetails, response, err := updateRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		return handleError(response, UPDATE, err)
	}

	//Read response
	model := currentModel.readAPIKeyDetails(apiKeyUserDetails)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update Completed",
		ResourceModel:   model}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(DeleteRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	deleteRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.DeleteApiKey(
		context.Background(),
		*currentModel.OrgId,
		*currentModel.Id,
	)
	_, response, err := deleteRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		return handleError(response, DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
		ResourceModel:   nil}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()

	modelValidation := validator.ValidateModel(ListRequiredFields, currentModel)
	if modelValidation != nil {
		return *modelValidation, nil
	}

	// Create atlas client
	if currentModel.Profile == nil || *currentModel.Profile == "" {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}

	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}
	apiKeyRequest := atlas.AtlasV2.ProgrammaticAPIKeysApi.ListApiKeys(
		context.Background(),
		*currentModel.OrgId,
	)

	pageNumber := 1
	itemsPerPage := 100
	apiKeyRequest = apiKeyRequest.ItemsPerPage(itemsPerPage).PageNum(pageNumber)

	pagedApiKeysList, response, err := apiKeyRequest.Execute()
	defer closeResponse(response)
	if err != nil {
		return handleError(response, LIST, err)
	}
	apiKeyList := pagedApiKeysList.Results

	defer closeResponse(response)
	for *pagedApiKeysList.TotalCount > len(apiKeyList) {
		pageNumber++
		apiKeyRequest = apiKeyRequest.PageNum(pageNumber)
		nextPageResults, response, err := apiKeyRequest.Execute()
		if err != nil {
			return handleError(response, LIST, err)
		}
		apiKeyList = append(apiKeyList, nextPageResults.Results...)
	}

	apiKeys := make([]interface{}, len(apiKeyList))
	for i := range apiKeyList {
		model := new(Model)
		model = model.readAPIKeyDetails(&apiKeyList[i])
		model.Profile = currentModel.Profile
		model.OrgId = currentModel.OrgId
		apiKeys[i] = model
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Completed",
		ResourceModels:  apiKeys}, nil
}

func closeResponse(response *http.Response) {
	if response != nil {
		response.Body.Close()
	}
}

func handleError(response *http.Response, method string, err error) (handler.ProgressEvent, error) {
	errMsg := fmt.Sprintf("%s error:%s", method, err.Error())
	_, _ = logger.Warn(errMsg)
	if response.StatusCode == http.StatusConflict {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: cloudformation.HandlerErrorCodeAlreadyExists}, nil
	}
	if response.StatusCode == http.StatusBadRequest {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          errMsg,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}
	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}

func (m *Model) readAPIKeyDetails(apikey *atlasSDK.ApiKeyUserDetails) *Model {
	model := new(Model)
	model.Profile = m.Profile
	model.OrgId = m.OrgId
	model.Id = apikey.Id
	model.Description = apikey.Desc
	model.PublicKey = apikey.PublicKey
	model.PrivateKey = apikey.PrivateKey
	var roles []string
	for i := range apikey.Roles {
		if apikey.Roles[i].RoleName != nil {
			roles = append(roles, *apikey.Roles[i].RoleName)
		}
	}
	model.Roles = roles

	var links []Link
	for i := range apikey.Links {
		link := Link{
			Href: apikey.Links[i].Href,
			Rel:  apikey.Links[i].Rel,
		}
		links = append(links, link)
	}
	model.Links = links
	return model
}
