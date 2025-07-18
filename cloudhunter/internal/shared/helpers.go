package shared

import (
	"encoding/json"
	"fmt"
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

func RenderBucketContent(nodes []*S3Node, indent string) {
	for _, node := range nodes {
		if node.IsFolder {
			fmt.Printf("%s %s\n", indent, node.Name)
			RenderBucketContent(node.Children, indent+"  ")
		} else {
			fmt.Printf("%s %s\n", indent, node.Name)
		}
	}
}
