package inspector

import (
	"fmt"
	"net/http"
	"time"
)

var securityHeaders = []string{
	"Content-Security-Policy",     // Prevents XSS by controlling allowed sources for scripts, styles, etc.
	"Strict-Transport-Security",   // Forces browsers to use HTTPS for future requests
	"X-Content-Type-Options",      // Prevents MIME-sniffing; enforces declared Content-Type
	"X-Frame-Options",             // Prevents clickjacking by controlling iframe embedding
	"Referrer-Policy",             // Controls how much referrer information is sent with requests
	"Permissions-Policy",          // Restricts use of browser features like camera, geolocation, etc.
	"Access-Control-Allow-Origin", // Defines allowed origins for CORS requests
	"Set-Cookie",                  // Used to check for Secure/HttpOnly/SameSite
}

func FetchHeaders(inputURL string) (map[string]string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// limiting redirec following to 3 levels
			if len(via) >= 3 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}

	req, err := http.NewRequest("GET", inputURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error making request to the URL: %v", err)
	}
	req.Header.Set("User-Agent", "HeaderInspector/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	defer resp.Body.Close()

	// collecting and checking headers
	results := make(map[string]string)
	for _, name := range securityHeaders {
		if values, ok := resp.Header[name]; ok && len(values) > 0 {

			if name == "Set-Cookie" {
				cookies := resp.Cookies()
				status := "✓ Present"
				for _, cookie := range cookies {
					if !cookie.Secure || !cookie.HttpOnly {
						status = "⚠ Insecure Flags"
						break
					}
				}
				results[name] = status
			} else {
				results[name] = "✓ Found"
			}
		} else {
			results[name] = "✗ Missing"
		}
	}
	return results, nil
}
