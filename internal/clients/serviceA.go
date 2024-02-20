package clients

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/iriojose/cada_test/pkg/aws"
	"github.com/iriojose/cada_test/pkg/types"
	"github.com/iriojose/cada_test/pkg/utils"
)

func ServiceA(pairs []string, response chan types.ResponseService, errorResponse chan error) {
	var serviceA types.ServiceAResponse
	URL := os.Getenv("SERVICE_A_URL")
	APIKEY, err := aws.GetApikey(os.Getenv("SECRET_NAME_SERVICE_A"))

	if err != nil {
		response <- serviceA
		errorResponse <- err
	}

	URL = fmt.Sprintf(`%s?api_key=%s&base=%s&target=%s`, URL, APIKEY, pairs[0], pairs[1])

	byteValue, err := utils.HttpGET(URL)

	if err != nil {
		response <- serviceA
		errorResponse <- err
	}

	json.Unmarshal(byteValue, &serviceA)

	if serviceA.Base == "" {
		response <- serviceA
		errorResponse <- fmt.Errorf("response invalid from service A")
	}

	response <- serviceA
	errorResponse <- nil
}
