// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package faker_test

import (
	"testing"

	"github.com/khulnasoft/faker"
	"github.com/stretchr/testify/assert"
)

func TestNames(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Names(-1) should panic()")
		}
	}()
	_ = faker.Names(-1)

	names := faker.Names(0)
	assert.NotNil(t, names)
	assert.Empty(t, names)

	max := 100
	names = faker.Names(max)
	assert.NotNil(t, names)
	assert.LessOrEqual(t, len(names), max)
	for _, n := range names {
		assert.NotEmpty(t, n)
	}
}
