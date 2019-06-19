package k8s

import (
	"fmt"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

const ResourceNodeName = "nodeName"

type Metrics struct{}

func (m *Metrics) NodeResourceList() ([]map[coreV1.ResourceName] interface{}, error) {
	var (
		err error
		nodeResources [] map[coreV1.ResourceName] interface{}
		nodeMetricsList *v1beta1.NodeMetricsList
		node *coreV1.Node
		cpuRequestsFraction, memoryRequestsFraction float64 = 0, 0
	)

	config, err := GetRestConf()
	if err != nil {
		return nodeResources, err
	}
	mc, err := metrics.NewForConfig(config)
	if err != nil {
		return nodeResources, err
	}
	nodeMetricsList, err = mc.MetricsV1beta1().NodeMetricses().List(metaV1.ListOptions{})
	if err != nil {
		return nodeResources, err
	}

	k8sNode := &Node{}

	for _, nodeMetric := range nodeMetricsList.Items {
		if node, err = k8sNode.GET(nodeMetric.Name); err != nil {
			return nodeResources, err
		}
		requestsFraction := make(map[coreV1.ResourceName] interface{})

		cpuRequests := nodeMetric.Usage.Cpu()
		memoryRequests := nodeMetric.Usage.Memory()

		memQuantity := node.Status.Allocatable[coreV1.ResourceMemory]
		totalMemAvail := int(memQuantity.Value() >> 20)
		fmt.Println(totalMemAvail)

		if capacity := float64(node.Status.Capacity.Cpu().MilliValue()); capacity > 0 {
			cpuRequestsFraction = float64(cpuRequests.MilliValue()) / capacity * 100
		}
		if capacity := float64(node.Status.Capacity.Memory().MilliValue()); capacity > 0 {
			memoryRequestsFraction = float64(memoryRequests.MilliValue()) / capacity * 100
		}
		requestsFraction[ResourceNodeName] = nodeMetric.Name
		requestsFraction[coreV1.ResourceCPU] = int64(cpuRequestsFraction)
		requestsFraction[coreV1.ResourceMemory] = int64(memoryRequestsFraction)

		nodeResources = append(nodeResources, requestsFraction)
	}
	return nodeResources, nil
}