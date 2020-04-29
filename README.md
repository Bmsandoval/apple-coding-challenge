## Setup
* First and foremost, this project was built using golang v1.13, ensure you are using this version
* Place this project _outside_ your go path
* Source the profile.sh in the project's root directory \
`$ . profile.sh`
* Install dependencies \
`$ go mod download`

## Development Procedure
* To run the project, execute the following from the project's root directory \
`$ go run .`
* If you update a service, you need to update the mocks. Run the following \
`$ applecc mock`

## Testing
* Mock testing is also built into the profile. \
`$ applecc test`
