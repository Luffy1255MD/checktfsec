package secrets

import (
	"testing"

	"github.com/aquasecurity/tfsec/internal/pkg/testutil"
)

func Test_AWSSensitiveVariables(t *testing.T) {
	expectedCode := "general-secrets-no-plaintext-exposure"

	var tests = []struct {
		name                  string
		source                string
		mustIncludeResultCode string
		mustExcludeResultCode string
	}{
		{
			name: "check sensitive variable with value",
			source: `
 variable "db_password" {
 	default = "something"
 }`,
			mustIncludeResultCode: expectedCode,
		},
		{
			name: "check sensitive variable without default",
			source: `
 variable "db_password" {
 
 }`,
			mustExcludeResultCode: expectedCode,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			results := testutil.ScanHCL(test.source, t)
			if test.mustIncludeResultCode != "" {
				testutil.AssertRuleFound(t, test.mustIncludeResultCode, results, "false negative found")
			}
			if test.mustExcludeResultCode != "" {
				testutil.AssertRuleNotFound(t, test.mustExcludeResultCode, results, "false positive found")
			}
		})
	}

}