// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of KhulnaSoft

package flow

import (
	"math/rand"

	flowpb "github.com/khulnasoft/shipyard/api/v1/flow"
)

// TrafficDirection returns randomly traffic direction INGRESS or EGRESS.
func TrafficDirection() flowpb.TrafficDirection {
	if rand.Intn(2) == 0 { // 50% chance of picking up ingress
		return flowpb.TrafficDirection_INGRESS
	}
	return flowpb.TrafficDirection_EGRESS
}
