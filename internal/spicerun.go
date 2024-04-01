package internal

import (
	"fmt"
	"net/url"
)

func ConnectSpiceTarget(address string, proxy string, insecure bool, vmid int, username string, token string, secret string, viewer_path string) error {
	authHeader := map[string]string{
		"Authorization": "PVEAPIToken=" + username + "!" + token + "=" + secret,
	}

	apiUrl, err := makeApiUrl(address)
	if err != nil {
		return fmt.Errorf("ERROR: Creating api url: %s", err)
	}

	_, err = makeProxyAddr(proxy, address)
	if err != nil {
		return fmt.Errorf("ERROR: Creating proxy address: %s", err)
	}

	if err := ensureAuthenticated(apiUrl, authHeader, insecure); err != nil {
		return fmt.Errorf("ERROR: Authentication failed: %s", err)
	}

	return nil
}

func makeApiUrl(address string) (string, error) {
	apiPath := "/api2/json"

	parsedUrl, err := url.Parse(address)
	if err != nil {
		return "", err
	}

	apiUrl := parsedUrl.Scheme + "://" + parsedUrl.Host + apiPath
	return apiUrl, nil
}

func makeProxyAddr(proxy string, hostAddress string) (string, error) {
	if proxy == "" {
		parsedUrl, err := url.Parse(hostAddress)
		if err != nil {
			return "", err
		}
		return parsedUrl.Hostname(), nil
	}
	return proxy, nil
}

func ensureAuthenticated(api_url string, authHeader map[string]string, insecure bool) error {
	_, err := sendApiRequest(api_url+"/access/permissions", "GET", authHeader, nil, insecure)
	return err
}
