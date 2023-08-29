#!/usr/bin/env bash

# Copyright 2023 MongoDB Inc
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -x
set -Eeou pipefail

AWS_SSM_Document_Name="gov-CFN-MongoDB-Atlas-Resource-Register"
BuilderRole="DevOpsIntegrationsContractors-CodeBuild"
AssumeRole="arn:aws:iam::711489243244:role/DevOpsIntegrationsContractorsSSM"
LogDeliveryBucket="atlascfnpublishing"
Repository="https://github.com/mongodb/mongodbatlas-cloudformation-resources"
BranchName="master"
ExecutionRoleName="DevOpsIntegrationsContractorsSSM"
Document_Region="us-east-1"
AccountIds="711489243244"
TargetLocationsMaxConcurrency="30"
Document_Version="\$DEFAULT"

RESOURCES="custom-db-role"
REGIONS="us-west-1"

if [ -z "${RESOURCES+x}" ];then
  echo "ATLAS_ORG_ID must be set"
  exit 1
fi

if [ -z "${REGIONS+x}" ];then
  echo "REGIONS must be set"
  exit 1
fi

echo "resources: ${RESOURCES}"
echo "regions: ${REGIONS}"

#  loop for RESOURCES
IFS=','
read -ra ResourceNames <<< "$RESOURCES"

# iterate over the array of resources
for ResourceName in "${ResourceNames[@]}"; do
    echo "resource: $ResourceName"

    # Currently OTHER_PARAMS is only used for
    # 1. federated-settings-org-role-mapping
    # 2. trigger
    # 3. ldap-configuration
    # 4. ldap-verify
    # 5. thirdpartyintegrations

    Path="cfn-resources/${ResourceName}/"
    CodeBuild_Project_Name="${ResourceName}-project-$((1 + RANDOM % 100))"



    jq --arg ExecutionRoleName "${ExecutionRoleName}" \
        --arg TargetLocationsMaxConcurrency "${TargetLocationsMaxConcurrency}" \
        --arg AccountIds "${AccountIds}" \
        --arg Regions "${REGIONS}" \
        '.[0].ExecutionRoleName?|=$ExecutionRoleName |
        .[0].TargetLocationMaxConcurrency?|=$TargetLocationsMaxConcurrency |
        .[0].Accounts[0]?|=$AccountIds |
        .[0].Regions?|=($Regions | gsub(" "; "") | split(",")) ' \
        "$(dirname "$0")/locations.json" >tmp.$$.json && mv tmp.$$.json "$(dirname "$0")/locations-temp.json"


    jq --arg Repository "${Repository}" \
       --arg ResourceName "${ResourceName}" \
       --arg OrgID "${ATLAS_ORG_ID}" \
       --arg PubKey "${ATLAS_PUBLIC_KEY}" \
       --arg PvtKey "${ATLAS_PRIVATE_KEY}" \
       --arg BranchName "${BranchName}" \
       --arg ProjectName "${CodeBuild_Project_Name}" \
       --arg OtherParams "${OTHER_PARAMS}" \
       --arg Path "${Path}" \
       --arg BuilderRole "${BuilderRole}" \
       --arg AssumeRole "${AssumeRole}" \
       --arg LogDeliveryBucket "${LogDeliveryBucket}" \
      '.Repository[0]?|=$Repository |
      .ResourceName[0]?|=$ResourceName |
      .OrgID[0]?|=$OrgID |
      .PubKey[0]?|=$PubKey |
      .PvtKey[0]?|=$PvtKey |
      .ProjectName[0]?|=$ProjectName |
      .BranchName[0]?|=$BranchName |
      .Path[0]?|=$Path |
      .BuilderRole[0]?|=$BuilderRole |
      .AssumeRole[0]?|=$AssumeRole |
      .LogDeliveryBucket[0]?|=$LogDeliveryBucket ' \
      "$(dirname "$0")/params.json" >tmp.$$.json && mv tmp.$$.json "$(dirname "$0")/params-temp.json"


    params_json_content=$(cat "$(dirname "$0")"/params-temp.json)
    locations_json_content=$(cat "$(dirname "$0")"/locations-temp.json)


    # use the aws cli to start the automation execution
    aws ssm start-automation-execution \
        --document-name  ${AWS_SSM_Document_Name}\
        --document-version ${Document_Version} \
        --parameters "${params_json_content}" \
        --target-locations "${locations_json_content}" \
        --region "${Document_Region}"

done

