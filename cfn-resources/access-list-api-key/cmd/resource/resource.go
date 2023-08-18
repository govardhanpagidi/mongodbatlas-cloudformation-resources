package resource

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/logger"
	progress_events "github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	atlasSDK "go.mongodb.org/atlas-sdk/v20230201002/admin"
)

var CreateRequiredFields = []string{constants.OrgID, constants.ApiUserID}
var ReadRequiredFields = []string{constants.OrgID, constants.ApiUserID, constants.IpAddress}
var DeleteRequiredFields = []string{constants.OrgID, constants.ApiUserID, constants.IpAddress}
var ListRequiredFields = []string{constants.OrgID, constants.IpAddress, constants.ApiUserID}

const (
	CREATE = "CREATE"
	READ   = "READ"
	DELETE = "DELETE"
	LIST   = "LIST"
)

// validateModel to validate inputs to all actions
func validateModel(fields []string, model *Model) *handler.ProgressEvent {
	return validator.ValidateModel(fields, model)
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if errEvent := validateModel(CreateRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	orgID := *currentModel.OrgId
	apiKeyID := *currentModel.ApiUserId

	if currentModel.CidrBlock == nil && currentModel.IpAddress == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Either IpAddress or CidrBlock is required",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if currentModel.CidrBlock != nil && currentModel.IpAddress != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "IpAddress and CidrBlock are mutually exclusive",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	// createReq.ApiService.
	entryList := make([]atlasSDK.UserAccessList, 0)
	var access atlasSDK.UserAccessList
	if currentModel.CidrBlock != nil {
		access.CidrBlock = currentModel.CidrBlock
	}
	if currentModel.IpAddress != nil {
		access.IpAddress = currentModel.IpAddress
	}
	entryList = append(entryList, access)

	createAccessListAPIKey := atlas.AtlasV2.ProgrammaticAPIKeysApi.CreateApiKeyAccessList(context.Background(), orgID, apiKeyID, &entryList)
	_, response, err := createAccessListAPIKey.Execute()

	defer closeResponse(response)
	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, CREATE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if errEvent := validateModel(ReadRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	orgID := *currentModel.OrgId
	apiKeyID := *currentModel.ApiUserId
	if currentModel.CidrBlock == nil && currentModel.IpAddress == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Either IpAddress or CidrBlock is required",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if currentModel.CidrBlock != nil && currentModel.IpAddress != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "IpAddress and CidrBlock are mutually exclusive",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	var access atlasSDK.UserAccessList
	if currentModel.CidrBlock != nil {
		access.CidrBlock = currentModel.CidrBlock
	}
	if currentModel.IpAddress != nil {
		access.IpAddress = currentModel.IpAddress
	}

	readAccessListAPIKey := atlas.AtlasV2.ProgrammaticAPIKeysApi.GetApiKeyAccessList(context.Background(), orgID, *access.IpAddress, apiKeyID)
	_, response, err := readAccessListAPIKey.Execute()
	defer closeResponse(response)
	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, READ, err)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Completed",
		ResourceModel:   currentModel}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: Update")
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if errEvent := validateModel(DeleteRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	orgID := *currentModel.OrgId
	apiKeyID := *currentModel.ApiUserId
	if currentModel.CidrBlock == nil && currentModel.IpAddress == nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "Either IpAddress or CidrBlock is required",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	if currentModel.CidrBlock != nil && currentModel.IpAddress != nil {
		return handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          "IpAddress and CidrBlock are mutually exclusive",
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, nil
	}

	var access atlasSDK.UserAccessList
	if currentModel.CidrBlock != nil {
		access.CidrBlock = currentModel.CidrBlock
	}
	if currentModel.IpAddress != nil {
		access.IpAddress = currentModel.IpAddress
	}

	deleteAccessListAPIKey := atlas.AtlasV2.ProgrammaticAPIKeysApi.DeleteApiKeyAccessListEntry(context.Background(), orgID, apiKeyID, *access.IpAddress)
	_, response, err := deleteAccessListAPIKey.Execute()

	defer closeResponse(response)
	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, DELETE, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Completed",
		ResourceModel:   nil}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if errEvent := validateModel(ListRequiredFields, currentModel); errEvent != nil {
		return *errEvent, nil
	}

	// Create atlas client
	if currentModel.Profile == nil {
		currentModel.Profile = aws.String(profile.DefaultProfile)
	}
	atlas, peErr := util.NewAtlasClient(&req, currentModel.Profile)
	if peErr != nil {
		return *peErr, nil
	}

	orgID := *currentModel.OrgId
	apiKeyID := *currentModel.ApiUserId

	listAccessListAPIKey := atlas.AtlasV2.ProgrammaticAPIKeysApi.ListApiKeyAccessListsEntries(context.Background(), orgID, apiKeyID)

	accessListResponse, response, err := listAccessListAPIKey.Execute()

	defer closeResponse(response)

	if err != nil {
		_, _ = logger.Warnf("Execute error: %s", err.Error())
		return handleError(response, LIST, err)
	}

	accessListModels := make([]interface{}, 0)
	for i := range accessListResponse.Results {
		l := accessListResponse.Results[i]
		label := Model{
			CidrBlock: l.CidrBlock,
			ApiUserId: currentModel.ApiUserId,
			OrgId:     currentModel.OrgId,
			Profile:   currentModel.Profile,
			IpAddress: l.IpAddress,
		}
		accessListModels = append(accessListModels, label)
	}
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List Complete",
		ResourceModels:  accessListModels,
	}, nil
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
	return progress_events.GetFailedEventByResponse(errMsg, response), nil
}