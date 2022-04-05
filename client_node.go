package themisclient

import (
	"context"
	"errors"

	themis "go.themis.run/themisclient/pb"
)

type NodeInfo struct {
	Name        string
	Address     string
	RaftAddress string
	Term        int32
	Role        string
	LogTerm     int32
	LogIndex    int32
}

var ErrorNodeNotFound = errors.New("node not found")

func (c *Client) Info() Info {
	return *c.info
}

func (c *Client) NodeInfo(name string) (NodeInfo, error) {
	addr, ok := c.info.Servers[name]
	if !ok {
		return NodeInfo{}, ErrorNodeNotFound
	}

	tclient, err := c.newClient(addr)
	if err != nil {
		return NodeInfo{}, err
	}

	resp, err := tclient.Info(context.Background(), &themis.InfoRequest{})
	if err != nil {
		return NodeInfo{}, err
	}

	return NodeInfo{
		Name:        resp.Name,
		Address:     resp.Address,
		RaftAddress: resp.RaftAddress,
		Term:        resp.Term,
		Role:        resp.Role,
		LogTerm:     resp.LogTerm,
		LogIndex:    resp.LogIndex,
	}, nil
}

func (c *Client) SearchKVListFromNodeName(name string) ([]*themis.KV, error) {
	addr, ok := c.info.Servers[name]
	if !ok {
		return nil, ErrorNodeNotFound
	}

	tclient, err := c.newClient(addr)
	if err != nil {
		return nil, err
	}

	resp, err := tclient.SearchByPrefix(context.Background(), &themis.SearchRequest{})
	if err != nil {
		return nil, err
	}
	return resp.KvList, nil
}
