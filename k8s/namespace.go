package k8s

import (
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Namespace struct{}

func (ns *Namespace) List() (*coreV1.NamespaceList, error) {
	var (
		clientSet     *kubernetes.Clientset
		namespaceList *coreV1.NamespaceList
		err           error
	)

	if clientSet, err = initClient(); err != nil {
		fmt.Println(err)
		return namespaceList, err
	}

	if namespaceList, err = clientSet.CoreV1().Namespaces().List(metaV1.ListOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			fmt.Println(err)
			return namespaceList, err
		}
	}
	return namespaceList, nil
}
