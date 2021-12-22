// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"
	"fmt"

	"github.com/Arman92/go-tdlib/v2/tdlib"
)

// GetStatisticalGraph Loads an asynchronous or a zoomed in statistical graph
// @param chatID Chat identifier
// @param token The token for graph loading
// @param x X-value for zoomed in graph or 0 otherwise
func (client *Client) GetStatisticalGraph(chatID int64, token string, x int64) (tdlib.StatisticalGraph, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":   "getStatisticalGraph",
		"chat_id": chatID,
		"token":   token,
		"x":       x,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	switch tdlib.StatisticalGraphEnum(result.Data["@type"].(string)) {

	case tdlib.StatisticalGraphDataType:
		var statisticalGraph tdlib.StatisticalGraphData
		err = json.Unmarshal(result.Raw, &statisticalGraph)
		return &statisticalGraph, err

	case tdlib.StatisticalGraphAsyncType:
		var statisticalGraph tdlib.StatisticalGraphAsync
		err = json.Unmarshal(result.Raw, &statisticalGraph)
		return &statisticalGraph, err

	case tdlib.StatisticalGraphErrorType:
		var statisticalGraph tdlib.StatisticalGraphError
		err = json.Unmarshal(result.Raw, &statisticalGraph)
		return &statisticalGraph, err

	default:
		return nil, fmt.Errorf("Invalid type")
	}
}
