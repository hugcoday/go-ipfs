package coreunix

import (
	"context"

	core "github.com/ipsn/go-ipfs/core"
	uio "github.com/ipsn/go-ipfs/unixfs/io"
	path "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-path"
	resolver "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-path/resolver"
)

func Cat(ctx context.Context, n *core.IpfsNode, pstr string) (uio.DagReader, error) {
	r := &resolver.Resolver{
		DAG:         n.DAG,
		ResolveOnce: uio.ResolveUnixfsOnce,
	}

	dagNode, err := core.Resolve(ctx, n.Namesys, r, path.Path(pstr))
	if err != nil {
		return nil, err
	}

	return uio.NewDagReader(ctx, dagNode, n.DAG)
}
