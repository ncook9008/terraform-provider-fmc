package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccessPolicyDefaultActionIntrusionPolicy struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type AccessPolicyDefaultActionSyslogConfig struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type AccessPolicyDefaultAction struct {
	Intrusionpolicy AccessPolicyDefaultActionIntrusionPolicy `json:"intrusionPolicy"`
	Syslogconfig    AccessPolicyDefaultActionSyslogConfig    `json:"syslogConfig"`
	Type            string                                   `json:"type"`
	Logbegin        string                                   `json:"logBegin"`
	Logend          string                                   `json:"logEnd"`
	Sendeventstofmc string                                   `json:"sendEventsToFMC"`
	Action          string                                   `json:"action"`
	// Variableset struct {
	// 	ID   string `json:"id"`
	// 	Type string `json:"type"`
	// } `json:"variableSet"`
	// Snmpconfig struct {
	// 	ID   string `json:"id"`
	// 	Type string `json:"type"`
	// } `json:"snmpConfig"`
}

type AccessPolicy struct {
	Type          string                    `json:"type"`
	Name          string                    `json:"name"`
	Description   string                    `json:"description"`
	Defaultaction AccessPolicyDefaultAction `json:"defaultAction"`
}

type AccessPolicyResponse struct {
	Type  string `json:"type"`
	Rules struct {
		Reftype string `json:"refType"`
		Type    string `json:"type"`
		Links   struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"rules"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

type AccessPoliciesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []struct {
		Links struct {
			Self   string `json:"self"`
			Parent string `json:"parent"`
		} `json:"links"`
		Type string `json:"type"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"items"`
	Paging struct {
		Offset int      `json:"offset"`
		Limit  int      `json:"limit"`
		Count  int      `json:"count"`
		Next   []string `json:"next"`
		Pages  int      `json:"pages"`
	} `json:"paging"`
}

func (v *Client) GetAccessPolicyByName(ctx context.Context, name string) (*AccessPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/accesspolicies?expanded=false&filter=name:%s", v.domainBaseURL, name)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting access policy by name/value: %s - %s", url, err.Error())
	}
	resp := &AccessPoliciesResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting access policy by name/value: %s - %s", url, err.Error())
	}
	switch l := len(resp.Items); {
	case l > 1:
		return nil, fmt.Errorf("duplicates found, length of response is: %d, expected 1, please search using a unique id, name or value", l)
	case l == 0:
		return nil, fmt.Errorf("no access policies found, length of response is: %d, expected 1, please check your filter", l)
	}
	return v.GetAccessPolicy(ctx, resp.Items[0].ID)
}

// /fmc_config/v1/domain/DomainUUID/policy/accesspolicies?bulk=true ( Bulk POST operation on access policies. )

func (v *Client) CreateAccessPolicy(ctx context.Context, accessPolicy *AccessPolicy) (*AccessPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/accesspolicies", v.domainBaseURL)
	body, err := json.Marshal(&accessPolicy)
	if err != nil {
		return nil, fmt.Errorf("creating access policies: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating access policies: %s - %s", url, err.Error())
	}
	item := &AccessPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating access policies: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetAccessPolicy(ctx context.Context, id string) (*AccessPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/accesspolicies/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting access policies: %s - %s", url, err.Error())
	}
	item := &AccessPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting access policies: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteAccessPolicy(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/policy/accesspolicies/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting access policies: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}