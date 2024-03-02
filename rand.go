// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package faker

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
