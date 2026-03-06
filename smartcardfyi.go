// Package smartcardfyi provides a Go client for the SmartCardFYI API.
//
// SmartCardFYI is a comprehensive smart card reference covering cards,
// platforms, standards, manufacturers, applications, form factors, and
// certifications. This client requires no authentication and has zero
// external dependencies.
//
// Usage:
//
//	client := smartcardfyi.NewClient()
//	result, err := client.Search("javacard")
package smartcardfyi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DefaultBaseURL is the default base URL for the SmartCardFYI API.
const DefaultBaseURL = "https://smartcardfyi.com/api"

// Client is a SmartCardFYI API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new SmartCardFYI API client with default settings.
func NewClient() *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) get(path string, result interface{}) error {
	resp, err := c.HTTPClient.Get(c.BaseURL + path)
	if err != nil {
		return fmt.Errorf("smartcardfyi: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("smartcardfyi: HTTP %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("smartcardfyi: decode failed: %w", err)
	}
	return nil
}

// Search searches across smart cards, standards, and glossary terms.
func (c *Client) Search(query string) (*SearchResult, error) {
	var result SearchResult
	path := fmt.Sprintf("/search/?q=%s", url.QueryEscape(query))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Card returns details for a smart card by slug.
func (c *Client) Card(slug string) (*CardDetail, error) {
	var result CardDetail
	if err := c.get("/card/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Platform returns details for a smart card platform by slug.
func (c *Client) Platform(slug string) (*PlatformDetail, error) {
	var result PlatformDetail
	if err := c.get("/platform/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Standard returns details for a smart card standard by slug.
func (c *Client) Standard(slug string) (*StandardDetail, error) {
	var result StandardDetail
	if err := c.get("/standard/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Manufacturer returns details for a smart card manufacturer by slug.
func (c *Client) Manufacturer(slug string) (*ManufacturerDetail, error) {
	var result ManufacturerDetail
	if err := c.get("/manufacturer/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Application returns details for a smart card application by slug.
func (c *Client) Application(slug string) (*ApplicationDetail, error) {
	var result ApplicationDetail
	if err := c.get("/application/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FormFactor returns details for a smart card form factor by slug.
func (c *Client) FormFactor(slug string) (*FormFactorDetail, error) {
	var result FormFactorDetail
	if err := c.get("/form-factor/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Certification returns details for a smart card certification by slug.
func (c *Client) Certification(slug string) (*CertificationDetail, error) {
	var result CertificationDetail
	if err := c.get("/certification/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GlossaryTerm returns a glossary term by slug.
func (c *Client) GlossaryTerm(slug string) (*GlossaryTerm, error) {
	var result GlossaryTerm
	if err := c.get("/glossary/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Compare compares two smart cards.
func (c *Client) Compare(slugA, slugB string) (*CompareResult, error) {
	var result CompareResult
	path := fmt.Sprintf("/compare/?a=%s&b=%s", url.QueryEscape(slugA), url.QueryEscape(slugB))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Random returns a random smart card.
func (c *Client) Random() (*CardDetail, error) {
	var result CardDetail
	if err := c.get("/random/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}
