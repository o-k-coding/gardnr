package mdparser

type MDNode struct {
	Value    string
	Type     string
	Children []*MDNode
	Parent   *MDNode
}

// TODO need to parse in an "iterable" way by feeding chunks of the file in and building the structure from that

// Super naive md parsing, does not enforce any rules, just packs
// the markdown into a basic tree structure.
// Only supports certain md elements currently
// some are just not relevant to how grdnr works yet and are just treated as text.
// Currently only supports parsing headers as nodes
// func Parse(markdown string, parent *MDNode) MDNode {
// 	if parent == nil {
// 		// Find the root. If not root exists, fail
// 	}
// 	// Otherwise find the support
// 	return nil
// }
