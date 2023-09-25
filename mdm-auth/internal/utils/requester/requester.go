package requester

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"mdm-auth/internal/utils/logging"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func Post(ctx context.Context, baseUrl string, payload url.Values) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(
		http.MethodPost,
		baseUrl,
		strings.NewReader(payload.Encode()),
	)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(payload.Encode())))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logging.Error("Body.Close  err:", err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if !IsOkCode(res.StatusCode) {
		var errResp ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, err
		}
		return nil, errors.New(errResp.ErrorDescription)
	}
	return body, nil
}

func IsOkCode(code int) bool {
	return (code >= http.StatusOK && code < http.StatusMultipleChoices) || code == 0
}

type ErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
