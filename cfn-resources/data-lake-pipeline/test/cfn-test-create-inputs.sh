#!/usr/bin/env bash
# cfn-test-create-inputs.sh
#
# This tool generates json files in the inputs/ for `cfn test`.
#

set -o errexit
set -o nounset
set -o pipefail

set -x

function usage {
    echo "usage:$0 <project/cluster_name>"
    echo "Creates a new Data lake pipeline with the cluster details for the test"
}

if [ "$#" -ne 2 ]; then usage; fi
if [[ "$*" == help ]]; then usage; fi

rm -rf inputs
mkdir inputs

projectName="${1}"
projectId=$(atlas projects list --output json | jq --arg NAME "${projectName}" -r '.results[] | select(.name==$NAME) | .id')
if [ -z "$projectId" ]; then
    projectId=$(atlas projects create "${projectName}" --output=json | jq -r '.id')

    echo -e "Created project \"${projectName}\" with id: ${projectId}\n"
else
    echo -e "FOUND project \"${projectName}\" with id: ${projectId}\n"
fi

echo "Check if a project is created $projectId"

region="us-east-1"

clusterName="${projectName}"

atlas clusters create "${clusterName}" --projectId "${projectId}" --backup --provider AWS --region US_EAST_1 --members 3 --tier M10 --mdbVersion 5.0 --diskSizeGB 10 --output=json
atlas clusters watch "${clusterName}" --projectId "${projectId}"
echo -e "Created Cluster \"${clusterName}\""

jq --arg region "$region" \
   --arg clusterName "$clusterName" \
   --arg projectId "$projectId" \
   '.Source.ClusterName?|=$clusterName | .ProjectId?|=$projectId ' \
   "$(dirname "$0")/inputs_1_create.template.json" > "inputs/inputs_1_create.json"

jq --arg region "$region" \
   --arg clusterName "$clusterName" \
   --arg projectId "$projectId" \
   '.Source.ClusterName?|=$clusterName | .ProjectId?|=$projectId ' \
   "$(dirname "$0")/inputs_1_update.template.json" > "inputs/inputs_1_update.json"

#SET INVALID NAME
clusterName="^%LKJ)(*J_ {+_+O_)"

jq --arg region "$region" \
   --arg clusterName "$clusterName" \
   --arg projectId "$projectId" \
   '.Source.ClusterName?|=$clusterName | .ProjectId?|=$projectId ' \
   "$(dirname "$0")/inputs_1_invalid.template.json" > "inputs/inputs_1_invalid.json"
