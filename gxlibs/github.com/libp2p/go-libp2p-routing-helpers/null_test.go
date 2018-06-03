package routinghelpers

import (
	"context"
	"testing"

	peer "github.com/ipsn/go-ipfs/gxlibs/github.com/libp2p/go-libp2p-peer"
	routing "github.com/ipsn/go-ipfs/gxlibs/github.com/libp2p/go-libp2p-routing"
)

func TestNull(t *testing.T) {
	var n Null
	ctx := context.Background()
	if err := n.PutValue(ctx, "anything", nil); err != routing.ErrNotSupported {
		t.Fatal(err)
	}
	if _, err := n.GetValue(ctx, "anything", nil); err != routing.ErrNotFound {
		t.Fatal(err)
	}
	if err := n.Provide(ctx, nil, false); err != routing.ErrNotSupported {
		t.Fatal(err)
	}
	if _, ok := <-n.FindProvidersAsync(ctx, nil, 10); ok {
		t.Fatal("expected no values")
	}
	if _, err := n.FindPeer(ctx, peer.ID("thing")); err != routing.ErrNotFound {
		t.Fatal(err)
	}
}