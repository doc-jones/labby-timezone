package main
import (
    "context"
    "encoding/json"
	"time"
	"net/http"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)
// response that supports Json serialization
type response struct {
    UTC time.Time `json:"utc"`
}
// request handler that creates a response with current time(UTC) serialize to Json
func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    now := time.Now()
    resp := &response{
        UTC: now.UTC(),
    }
    body, err := json.Marshal(resp)
    if err != nil {
        return events.APIGatewayProxyResponse{}, err
    }
    return events.APIGatewayProxyResponse{Body: string(body), StatusCode: 200}, nil
}


// register handler using AWS Lambda for Go lib
func main() {
    lambda.Start(handleRequest)
}

// return time in local timezone based upon IP address
var httpClient = &http.Client{}
func timezone(ip string) *time.Location {
        resp, err := httpClient.Get("https://ipapi.co/" + ip + "/timezone/")
        if err != nil {
                return nil
        }
        defer resp.Body.Close()
        tz, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                return nil
        }
        loc, err := time.LoadLocation(string(tz))
        if err != nil {
                return nil
        }
		return loc
}

