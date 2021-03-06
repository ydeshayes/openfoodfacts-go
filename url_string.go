// Copyright © 2016 OpenFoodFacts. All rights reserved.
// Use of this source code is governed by the MIT license which can be found in the LICENSE.txt file.

package openfoodfacts

import (
	"fmt"
	"net/url"
	"strings"
)

// UrlString allows URL to be properly un/marshaled from a JSON string.
type UrlString struct {
	url.URL
}

func (u *UrlString) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return nil
	}

	parsed, err := url.Parse(s)
	if err != nil {
		return err
	}

	u.URL = *parsed
	return nil
}

func (u *UrlString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", u.URL.String())), nil
}
