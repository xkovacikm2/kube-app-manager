package kube

import (
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func Test_mapWebEndpointList_filtersAndMaps(t *testing.T) {
	resourceList := []unstructured.Unstructured{
		{Object: map[string]any{"spec": map[string]any{
			"name":        "  app-a  ",
			"description": " App A ",
			"url":         " https://a.example.com ",
			"icon":        " https://a.example.com/icon.svg ",
		}}},
		{Object: map[string]any{"spec": map[string]any{
			"description": "missing name",
		}}},
		{Object: map[string]any{"spec": map[string]any{
			"name": 123.0,
		}}},
		{Object: map[string]any{"data": map[string]any{}}},
	}

	applicationList := mapWebEndpointList(resourceList)
	if len(applicationList) != 1 {
		t.Fatalf("expected 1 application, got %d", len(applicationList))
	}

	if applicationList[0].Name != "app-a" {
		t.Fatalf("unexpected name: %s", applicationList[0].Name)
	}
	if applicationList[0].Description != "App A" {
		t.Fatalf("unexpected description: %s", applicationList[0].Description)
	}
	if applicationList[0].URL != "https://a.example.com" {
		t.Fatalf("unexpected url: %s", applicationList[0].URL)
	}
	if applicationList[0].Icon != "https://a.example.com/icon.svg" {
		t.Fatalf("unexpected icon: %s", applicationList[0].Icon)
	}
}
