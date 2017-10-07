// +build pro ent

package nomad

import (
	"fmt"

	memdb "github.com/hashicorp/go-memdb"
	"github.com/hashicorp/nomad/nomad/state"
	"github.com/hashicorp/nomad/nomad/structs"
)

var (
	// proContexts are the pro contexts which are searched to find matches
	// for a given prefix
	proContexts = []structs.Context{structs.Namespaces}
)

// getProMatch is used to match on an object only defined in Nomad Pro or
// Premium
func getProMatch(match interface{}) (id string, ok bool) {
	switch match.(type) {
	case *structs.Namespace:
		return match.(*structs.Namespace).Name, true
	default:
		return "", false
	}
}

// getProResourceIter is used to retrieve an iterator over an enterprise
// only table.
func getProResourceIter(context structs.Context, namespace, prefix string, ws memdb.WatchSet, state *state.StateStore) (memdb.ResultIterator, error) {
	switch context {
	case structs.Namespaces:
		return state.NamespacesByNamePrefix(ws, prefix)
	default:
		return nil, fmt.Errorf("context must be one of %v or 'all' for all contexts; got %q", allContexts, context)
	}
}
