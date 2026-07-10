package kube

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/xkovacikm2/kube-app-manager/internal/apps"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var webEndpointGroupVersionResource = schema.GroupVersionResource{
	Group:    "apps.kovko.top",
	Version:  "v1alpha1",
	Resource: "webendpoints",
}

type WebEndpointSource struct {
	dynamicClient dynamic.Interface
}

func NewWebEndpointSource() (*WebEndpointSource, error) {
	restConfiguration, err := buildKubernetesConfiguration()
	if err != nil {
		return nil, fmt.Errorf("build kubernetes config: %w", err)
	}

	dynamicClient, err := dynamic.NewForConfig(restConfiguration)
	if err != nil {
		return nil, fmt.Errorf("create dynamic client: %w", err)
	}

	return &WebEndpointSource{dynamicClient: dynamicClient}, nil
}

func (source *WebEndpointSource) ListApplications(context context.Context) ([]apps.Application, error) {
	resourceList, err := source.dynamicClient.Resource(webEndpointGroupVersionResource).List(context, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("list webendpoints: %w", err)
	}

	applicationList := mapWebEndpointList(resourceList.Items)
	sort.Slice(applicationList, func(indexOne int, indexTwo int) bool {
		return applicationList[indexOne].Name < applicationList[indexTwo].Name
	})

	return applicationList, nil
}

func mapWebEndpointList(resourceList []unstructured.Unstructured) []apps.Application {
	applicationList := make([]apps.Application, 0, len(resourceList))

	for _, resource := range resourceList {
		specObject, found, _ := unstructured.NestedMap(resource.Object, "spec")
		if !found {
			continue
		}

		name := readSpecString(specObject, "name")
		if name == "" {
			continue
		}

		applicationList = append(applicationList, apps.Application{
			Name:        name,
			Description: readSpecString(specObject, "description"),
			URL:         readSpecString(specObject, "url"),
			Icon:        readSpecString(specObject, "icon"),
		})
	}

	return applicationList
}

func readSpecString(specObject map[string]any, key string) string {
	value, found := specObject[key]
	if !found {
		return ""
	}

	stringValue, ok := value.(string)
	if !ok {
		return ""
	}

	return strings.TrimSpace(stringValue)
}

func buildKubernetesConfiguration() (*rest.Config, error) {
	inClusterConfiguration, err := rest.InClusterConfig()
	if err == nil {
		return inClusterConfiguration, nil
	}

	kubeConfigPath := strings.TrimSpace(os.Getenv("KUBECONFIG"))
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	if kubeConfigPath != "" {
		loadingRules.ExplicitPath = kubeConfigPath
	}

	clientConfiguration := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		&clientcmd.ConfigOverrides{},
	)

	restConfiguration, configErr := clientConfiguration.ClientConfig()
	if configErr != nil {
		return nil, fmt.Errorf("load kubeconfig: %w", configErr)
	}

	return restConfiguration, nil
}
