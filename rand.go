// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of KhulnaSoft

package fake

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
