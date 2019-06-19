package k8s

import (
	"fmt"
	"io/ioutil"
	appsV1Beta1 "k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

type Deployment struct{}

func (dp *Deployment) List() (*appsV1Beta1.DeploymentList, error) {
	var (
		clientSet      *kubernetes.Clientset
		deploymentList *appsV1Beta1.DeploymentList
		err            error
	)

	if clientSet, err = initClient(); err != nil {
		fmt.Println(err)
		return deploymentList, err
	}

	if deploymentList, err = clientSet.AppsV1beta1().Deployments("default").List(metaV1.ListOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			fmt.Println(err)
			return deploymentList, err
		}
	}
	return deploymentList, nil
}

func (dp *Deployment) GET(deploymentName string) (*appsV1Beta1.Deployment, error) {
	var (
		clientSet  *kubernetes.Clientset
		deployment *appsV1Beta1.Deployment
		err        error
	)

	if clientSet, err = initClient(); err != nil {
		return deployment, err
	}

	if deployment, err = clientSet.AppsV1beta1().Deployments("default").Get(deploymentName, metaV1.GetOptions{}); err != nil {
		return deployment, err
	}
	return deployment, nil
}

func (dp *Deployment) Create(deploymentName string) error {
	var (
		clientSet  *kubernetes.Clientset
		deployYaml []byte
		deployJson []byte
		deployment = appsV1Beta1.Deployment{}
		replicas   int32
		err        error
	)

	if clientSet, err = initClient(); err != nil {
		return err
	}

	if deployYaml, err = ioutil.ReadFile("./nginx.yaml"); err != nil {
		return err
	}

	if deployJson, err = yaml2.ToJSON(deployYaml); err != nil {
		return err
	}

	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		return err
	}

	replicas = 2
	deployment.Spec.Replicas = &replicas

	if _, err = clientSet.AppsV1beta1().Deployments("default").Get(deployment.Name, metaV1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		if _, err = clientSet.AppsV1beta1().Deployments("default").Create(&deployment); err != nil {
			return err
		}
	} else {
		if _, err = clientSet.AppsV1beta1().Deployments("default").Update(&deployment); err != nil {
			return err
		}
	}
	return nil
}

func (dp *Deployment) Delete(deploymentName string) error {
	var (
		clientSet *kubernetes.Clientset
		err       error
	)

	if clientSet, err = initClient(); err != nil {
		fmt.Println(err)
		return nil
	}

	if err = clientSet.AppsV1beta1().Deployments("default").Delete(deploymentName, &metaV1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func (dp *Deployment) Update(version string) error {
	var (
		clientSet  *kubernetes.Clientset
		deployYaml []byte
		deployJson []byte
		deployment = appsV1Beta1.Deployment{}
		replicas   int32
		err        error
	)

	if clientSet, err = initClient(); err != nil {
		return err
	}

	if deployYaml, err = ioutil.ReadFile("./nginx.yaml"); err != nil {
		return err
	}

	if deployJson, err = yaml2.ToJSON(deployYaml); err != nil {
		return err
	}

	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		return err
	}

	replicas = 2
	deployment.Spec.Replicas = &replicas

	if _, err = clientSet.AppsV1beta1().Deployments("default").Get(deployment.Name, metaV1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		if _, err = clientSet.AppsV1beta1().Deployments("default").Create(&deployment); err != nil {
			return err
		}
	} else {
		if _, err = clientSet.AppsV1beta1().Deployments("default").Update(&deployment); err != nil {
			return err
		}
	}
	return nil
}
