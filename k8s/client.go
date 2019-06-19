package k8s

import (
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// 获取k8s restful client配置
func GetRestConf() (restConf *rest.Config, err error) {
	var (
		kubeConfig []byte
	)

	// 读kubeConfig文件
	if kubeConfig, err = ioutil.ReadFile("./admin.conf"); err != nil {
		return nil, err
	}
	// 生成rest client配置
	if restConf, err = clientcmd.RESTConfigFromKubeConfig(kubeConfig); err != nil {
		return nil, err
	}
	return
}

// 初始化k8s客户端
func initClient() (clientSet *kubernetes.Clientset, err error) {
	var (
		restConf *rest.Config
	)

	if restConf, err = GetRestConf(); err != nil {
		return nil, err
	}

	// 生成clientSet配置
	if clientSet, err = kubernetes.NewForConfig(restConf); err != nil {
		return nil, err
	}
	return
}
