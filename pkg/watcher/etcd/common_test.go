package etcd

import (
	"testing"

	"github.com/coreos/etcd/clientv3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
)

func TestNoopKVClient(t *testing.T) {
	c := noopKVClient{}

	pr, err := c.Put(nil, "", "")
	assert.Nil(t, pr)
	assert.NoError(t, err)

	gr, err := c.Get(nil, "")
	assert.Nil(t, gr)
	assert.NoError(t, err)

	dr, err := c.Delete(nil, "")
	assert.Nil(t, dr)
	assert.NoError(t, err)

	cr, err := c.Compact(nil, 0)
	assert.Nil(t, cr)
	assert.NoError(t, err)

	or, err := c.Do(nil, clientv3.Op{})
	assert.Equal(t, clientv3.OpResponse{}, or)
	assert.NoError(t, err)

	tx := c.Txn(nil)
	assert.Nil(t, tx)
}

type kvClientMock struct {
	mock.Mock
	valuesPut map[string]string
}

func givenKVClientMock() *kvClientMock {
	return &kvClientMock{
		valuesPut: make(map[string]string),
	}
}

func (m *kvClientMock) Put(ctx context.Context, key, val string, opts ...clientv3.OpOption) (
	*clientv3.PutResponse, error) {
	args := m.Called(ctx, key, val, opts)
	m.valuesPut[key] = val
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*clientv3.PutResponse), args.Error(1)
}

func (m *kvClientMock) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (
	*clientv3.GetResponse, error) {
	args := m.Called(ctx, key, opts)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*clientv3.GetResponse), args.Error(1)
}

func (m *kvClientMock) Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (
	*clientv3.DeleteResponse, error) {
	args := m.Called(ctx, key, opts)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*clientv3.DeleteResponse), args.Error(1)
}

func (m *kvClientMock) Compact(ctx context.Context, rev int64, opts ...clientv3.CompactOption) (
	*clientv3.CompactResponse, error) {
	args := m.Called(ctx, rev, opts)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*clientv3.CompactResponse), args.Error(1)
}

func (m *kvClientMock) Do(ctx context.Context, op clientv3.Op) (clientv3.OpResponse, error) {
	args := m.Called(ctx, op)
	return args.Get(0).(clientv3.OpResponse), args.Error(1)
}

func (m *kvClientMock) Txn(ctx context.Context) clientv3.Txn {
	args := m.Called(ctx)
	return args.Get(0).(clientv3.Txn)
}

func (m *kvClientMock) onPut(key string, val interface{}) *mock.Call {
	return m.On(
		"Put",
		mock.AnythingOfType("*context.timerCtx"),
		key,
		val,
		mock.AnythingOfType("[]clientv3.OpOption"),
	)
}

func (m *kvClientMock) onDelete(key string) *mock.Call {
	return m.On(
		"Delete",
		mock.AnythingOfType("*context.timerCtx"),
		key,
		mock.AnythingOfType("[]clientv3.OpOption"),
	)
}