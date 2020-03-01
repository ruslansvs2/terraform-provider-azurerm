package validate

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func CdnEndpointDeliveryPolicyRuleName() schema.SchemaValidateFunc {
	return validation.StringMatch(
		regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]*$"),
		"The Delivery Policy Rule Name must start with a letter any may only contain letters and numbers.",
	)
}

func RuleActionUrlRedirectPath() schema.SchemaValidateFunc {
	return validation.StringMatch(
		regexp.MustCompile("^(/.*)?$"),
		"The Url Redirect Path must start with a slash.",
	)
}

func RuleActionUrlRedirectQueryString() schema.SchemaValidateFunc {
	return func(i interface{}, s string) ([]string, []error) {
		querystring := i.(string)

		if len(querystring) > 100 {
			return nil, []error{fmt.Errorf("The Url Query String's max length is 100.")}
		}

		re, _ := regexp.Compile("^[?&]")
		if re.MatchString(querystring) {
			return nil, []error{fmt.Errorf("The Url Query String must not start with a question mark or ampersand.")}
		}

		kvre, _ := regexp.Compile("^[^?&]+=[^?&]+$")
		kvs := strings.Split(querystring, "&")
		for _, kv := range kvs {
			if len(kv) > 0 && !kvre.MatchString(kv) {
				return nil, []error{fmt.Errorf("The Url Query String must be in <key>=<value> format and separated by an ampersand.")}
			}
		}

		return nil, nil
	}
}

func RuleActionUrlRedirectFragment() schema.SchemaValidateFunc {
	return validation.StringMatch(
		regexp.MustCompile("^([^#].*)?$"),
		"The Url Redirect Path must start with a slash.",
	)
}
