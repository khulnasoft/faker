// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Khulnasoft

package flow

import (
	"math/rand"

	flowpb "github.com/khulnasoft/shipyard/api/v1/flow"
	"github.com/khulnasoft/shipyard/pkg/monitor/api"
)

var allEventTypes = []int32{
	api.MessageTypeUnspec,
	api.MessageTypeDrop,
	api.MessageTypeDebug,
	api.MessageTypeCapture,
	api.MessageTypeTrace,
	api.MessageTypePolicyVerdict,
	api.MessageTypeAccessLog,
	api.MessageTypeAgent,
}

// EventType generates a random EventType.
func EventType() *flowpb.KhulnasoftEventType {
	typ := allEventTypes[rand.Intn(len(allEventTypes))]
	if typ == api.MessageTypeUnspec {
		return nil
	}
	return &flowpb.KhulnasoftEventType{
		Type: typ,
		// NOTE: AgentNotify* are the most numerous.
		SubType: int32(rand.Intn(13)),
	}
}
