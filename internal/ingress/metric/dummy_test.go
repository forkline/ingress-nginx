package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/util/sets"
)

func TestNewDummyCollector(t *testing.T) {
	c := NewDummyCollector()
	assert.NotNil(t, c)
	_, ok := c.(*DummyCollector)
	assert.True(t, ok)
}

func TestDummyCollectorNoPanic(t *testing.T) { //nolint:revive // t is needed for test framework
	c := NewDummyCollector()

	c.ConfigSuccess(123, true)
	c.SetAdmissionMetrics(1, 2, 3, 4, 5, 6)
	c.IncReloadCount()
	c.IncReloadErrorCount()
	c.IncOrphanIngress("ns", "ing", "type")
	c.DecOrphanIngress("ns", "ing", "type")
	c.IncCheckCount("ns", "ing")
	c.IncCheckErrorCount("ns", "ing")
	c.RemoveMetrics(nil, nil)
	c.Start("")
	c.Stop("")
	c.SetSSLInfo(nil)
	c.SetSSLExpireTime(nil)
	c.SetHosts(sets.Set[string]{})
	c.OnStartedLeading("leader")
	c.OnStoppedLeading("leader")
}
