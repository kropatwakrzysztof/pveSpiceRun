package internal

import (
	"fmt"
	"net/url"
)

func ConnectSpiceTarget(address string, proxy string, insecure bool, vmid int, username string, token string, secret string, viewerPath string) error {
	authHeader := map[string]string{
		"Authorization": "PVEAPIToken=" + username + "!" + token + "=" + secret,
	}

	apiUrl, err := makeApiUrl(address)
	if err != nil {
		return fmt.Errorf("ERROR: Creating api url: %s", err)
	}

	proxyAddr, err := makeProxyAddr(proxy, address)
	if err != nil {
		return fmt.Errorf("ERROR: Creating proxy address: %s", err)
	}

	if err := ensureAuthenticated(apiUrl, authHeader, insecure); err != nil {
		return fmt.Errorf("ERROR: Authentication failed: %s", err)
	}

	nodeName, vmType, err := getVMInfo(apiUrl, authHeader, insecure, vmid)
	if err != nil {
		return fmt.Errorf("ERROR: Getting vm info: %s", err)
	}

	if err := ensureVMStarted(apiUrl, authHeader, insecure, nodeName, vmType, vmid); err != nil {
		return fmt.Errorf("ERROR: Trying to start VM: %s", err)
	}

	connFilePath, err := generateRemoteViewerConnectionFile(apiUrl, authHeader, insecure, nodeName, vmType, vmid, proxyAddr)
	if err != nil {
		return fmt.Errorf("ERROR: Generating VirtViewer connection file: %s", err)
	}

	if err := startRemoteViewer(viewerPath, connFilePath); err != nil {
		return fmt.Errorf("ERROR: Trying to start VirtViewer: %s", err)
	}

	// every thing OK returning no error
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
	_, err := sendApiRequest("GET", api_url+"/access/permissions", authHeader, nil, insecure)
	return err
}

func getVMInfo(apiUrl string, authHeader map[string]string, insecure bool, vmid int) (string, string, error) {
	panic("unimplemented")
}

func ensureVMStarted(apiUrl string, authHeader map[string]string, insecure bool, nodeName string, vmType string, vmid int) error {
	panic("unimplemented")
}

func generateRemoteViewerConnectionFile(apiUrl string, authHeader map[string]string, insecure bool, nodeName string, vmType string, vmid int, proxyAddr string) (string, error) {
	panic("unimplemented")
}

func startRemoteViewer(viewerPath, connFilePath string) error {
	panic("unimplemented")
}
