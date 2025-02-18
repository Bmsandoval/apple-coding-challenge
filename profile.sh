#!/usr/bin/env bash

# Just gets the top level directory of this project. Useful for scripting within the project via relative file paths
APPLE_CC_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

applecc () {
    # if no command given force help page
    local OPTION
	if [[ "$1" != "" ]]; then
        OPTION=$1
    else
        OPTION="help"
    fi
	# handle input options
    case "${OPTION}" in
        'help')
echo "Usage: $ ${FUNCNAME} [option] [flags]
Options:
- help: show this menu
- mock: Mock all services (req: gomock)
- test: run all mock tests
"
        ;;
        'mock')
            appleccMockServices
        ;;
        'test')
          appleccTestServers
        ;;
        *)
            echo -e "ERROR: invalid option. Try..\n$ ${FUNCNAME} help"
        ;;
    esac
}

appleccTestServers () {
     go test $(go list $APPLE_CC_DIR/...)
}

# Generate mock files for all services, putting the results in the proper file. renames some stuff for consistency.
# If you update any services, recommend running this function to update the services for the tests.
appleccMockServices () {
    MOCK_FOLDER="services"
    SERVICE_DIR="${APPLE_CC_DIR}/${MOCK_FOLDER}"
    SERVICES=$(find "${SERVICE_DIR}" -maxdepth 1 -mindepth 1 -type d)
    for SERVICE_PATH in ${SERVICES}
    do
        SERVICE_FOLDER_NAME=$(basename "${SERVICE_PATH}")
        if [[ "${SERVICE_FOLDER_NAME}" == grpc* ]]; then
          PROTOS=$(find "${SERVICE_PATH}" | grep ".pb.go")
          for PROTO in ${PROTOS}
          do
            PROTO_FILE_NAME=$(basename "${PROTO}")
            PROTO_FILE_NAME_STRIP_EXT=${PROTO_FILE_NAME/.go/}
            PROTO_NO_PB=${PROTO_FILE_NAME_STRIP_EXT/.pb/}
            PROTO_REPLACED_NAME=${PROTO_FILE_NAME_STRIP_EXT/./_}
            PROTO_SERVICE_NAME=$(echo "${PROTO_NO_PB}" | awk '{for (i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) substr($i,2)} 1')
            mockgen \
                -package=${MOCK_FOLDER}_mocks \
                -destination=mocks/${MOCK_FOLDER}_mocks/"${PROTO_REPLACED_NAME}"_mock.go \
                github.com/bmsandoval/apple-coding-challenge/services/"${SERVICE_FOLDER_NAME}" "${PROTO_SERVICE_NAME}Client"
          done
        elif [[ -f ${SERVICE_PATH}/interface.go ]]; then
            FOLDER_NAME="${SERVICE_PATH##*/}"
            mockgen \
                -source="${SERVICE_PATH}"/interface.go \
                -destination=mocks/${MOCK_FOLDER}_mocks/"${FOLDER_NAME}"_mock.go \
                -package=${MOCK_FOLDER}_mocks \
                -mock_names Service=Mock_"${FOLDER_NAME}"
        fi
    done
}

# Check if a command exists in the environment
# Returns 0 if command found
package-installed () {
	result=$(compgen -A function -abck | grep "^$1$")
    # Note that in bash, non-zero exit codes are error codes. returning 0 means success
	if [[ "${result}" == "$1" ]]; then
		# package installed
		return 0
	else
		# package not installed
		return 1
	fi
}
