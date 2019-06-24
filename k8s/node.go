package k8s

import (
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Node struct {
}

func (n *Node) GET(nodeName string) (*coreV1.Node, error) {
	var (
		clientSet *kubernetes.Clientset
		err       error
		node      *coreV1.Node
	)

	if clientSet, err = initClient(); err != nil {
		return node, err
	}

	if node, err = clientSet.CoreV1().Nodes().Get(nodeName, metaV1.GetOptions{}); err != nil {
		return node, err
	}
	return node, nil
}
