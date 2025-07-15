package shared

import (
	"encoding/json"
	"log"
	"net/url"
)

func ParseJsonPolicyDocument(policyData string) string {
	decodedPolicy, err := url.QueryUnescape(policyData)
	if err != nil {
		log.Fatal(err)
	}

	var policyObj any
	err = json.Unmarshal([]byte(decodedPolicy), &policyObj)
	if err != nil {
		log.Fatal(err)
	}

	policy, err := json.MarshalIndent(policyObj, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(policy)
}
