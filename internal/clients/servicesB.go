package clients

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/iriojose/cada_test/pkg/aws"
	"github.com/iriojose/cada_test/pkg/types"
	"github.com/iriojose/cada_test/pkg/utils"
)

func ServiceB(pairs []string, response chan types.ResponseService, errorResponse chan error) {
	var serviceB types.ServiceBResponse
	URL := os.Getenv("SERVICE_B_URL")
	APIKEY, err := aws.GetApikey("serviceB")

	if err != nil {
		response <- serviceB
		errorResponse <- err
	}

	URL = fmt.Sprintf(`%s%s/pair/%s/%s`, URL, APIKEY, pairs[0], pairs[1])
	byteValue, err := utils.HttpGET(URL)

	if err != nil {
		response <- serviceB
		errorResponse <- err
	}

	json.Unmarshal(byteValue, &serviceB)

	if serviceB.Result == "" {
		response <- serviceB
		errorResponse <- fmt.Errorf("response invalid from service B")
	}

	response <- serviceB
	errorResponse <- nil
}
