package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"jv/team-tone-tuner/config"
	"net/http"
)

func GetGithubAccessToken(code string) (string, error) {
	clientID, clientSecret := config.LoadedConfig.Github.ClientId, config.LoadedConfig.Github.ClientSecret

	requestBodyMap := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
	}

	requestJSON, _ := json.Marshal(requestBodyMap)

	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(requestJSON))

	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	respBody, _ := io.ReadAll(resp.Body)

	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
	}

	var ghResp githubAccessTokenResponse

	err = json.Unmarshal(respBody, &ghResp)

	if err != nil {
		return "", err
	}

	return ghResp.AccessToken, nil
}

func GetGithubUsername(accessToken string) (string, error) {
	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)

	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))

	// Make the request
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	// Read the response as a byte slice
	respbody, _ := io.ReadAll(resp.Body)

	type GithubReturn struct {
		Login string `json:"login"`
	}

	var githubReturn GithubReturn

	_ = json.Unmarshal(respbody, &githubReturn)

	// Convert byte slice to string and return
	return githubReturn.Login, nil
}
