# Math-Web-API
Github web service API that will have endpoints for simple math arithmetic

# Prerequisites

## Go
Must have Go version 1.18 installed
* See here for more information on how to download and install - https://go.dev/dl/

## Postman Collection
* You can use this preset postman collection to make the calls to the API once the server is running
https://www.getpostman.com/collections/56921d2e123a15f816d1
    - Localhost URL is default {{Dev API URL}} at http://localhost:3000
    - API url is /api/v1/
    - Must set a header Accept: application/json
    - 5 API Calls can take in a JSON Body which can look like:
        - ```
            {
                "array": "1, 2, 3, 4, 5, 3,5,191,04,3,201,650,1,2385,3,21004,85,2012,5832,0205",
                "quantifier": "25"
            }
          ```
        - 5 API Calls:
            1. /api/v1/min
                - both fields required
            2. /api/v1/max
                - both fields required
            3. /api/v1/avg
                - quantifier not needed
            4. /api/v1/median
                - quantifier not needed
            5. /api/v1/percentile
                - both fields required
    - Array only accepts int64 types (for now)

## Starting application
Run ```go run main.go ``` in order to start fiber server. Server will haver server information in terminal. 