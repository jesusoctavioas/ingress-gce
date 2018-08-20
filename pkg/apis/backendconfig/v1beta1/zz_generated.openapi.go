// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.BackendConfig": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Properties: map[string]spec.Schema{
						"kind": {
							SchemaProps: spec.SchemaProps{
								Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"apiVersion": {
							SchemaProps: spec.SchemaProps{
								Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"metadata": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
							},
						},
						"spec": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.BackendConfigSpec"),
							},
						},
						"status": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.BackendConfigStatus"),
							},
						},
					},
				},
			},
			Dependencies: []string{
				"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.BackendConfigSpec", "k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.BackendConfigStatus"},
		},
		"k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.BackendConfigSpec": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "BackendConfigSpec is the spec for a BackendConfig resource",
					Properties: map[string]spec.Schema{
						"iap": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.IAPConfig"),
							},
						},
						"cdn": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.CDNConfig"),
							},
						},
						"securityPolicy": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.SecurityPolicyConfig"),
							},
						},
					},
				},
			},
			Dependencies: []string{
				"k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.CDNConfig", "k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.IAPConfig", "k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.SecurityPolicyConfig"},
		},
		"k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.CDNConfig": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "CDNConfig contains configuration for CDN-enabled backends.",
					Properties: map[string]spec.Schema{
						"enabled": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"boolean"},
								Format: "",
							},
						},
						"cachePolicy": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.CacheKeyPolicy"),
							},
						},
					},
					Required: []string{"enabled"},
				},
			},
			Dependencies: []string{
				"k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.CacheKeyPolicy"},
		},
		"k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.CacheKeyPolicy": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "CacheKeyPolicy contains configuration for how requests to a CDN-enabled backend are cached.",
					Properties: map[string]spec.Schema{
						"includeHost": {
							SchemaProps: spec.SchemaProps{
								Description: "If true, requests to different hosts will be cached separately.",
								Type:        []string{"boolean"},
								Format:      "",
							},
						},
						"includeProtocol": {
							SchemaProps: spec.SchemaProps{
								Description: "If true, http and https requests will be cached separately.",
								Type:        []string{"boolean"},
								Format:      "",
							},
						},
						"includeQueryString": {
							SchemaProps: spec.SchemaProps{
								Description: "If true, query string parameters are included in the cache key according to QueryStringBlacklist and QueryStringWhitelist. If neither is set, the entire query string is included and if false the entire query string is excluded.",
								Type:        []string{"boolean"},
								Format:      "",
							},
						},
						"queryStringBlacklist": {
							SchemaProps: spec.SchemaProps{
								Description: "Names of query strint parameters to exclude from cache keys. All other parameters are included. Either specify QueryStringBlacklist or QueryStringWhitelist, but not both.",
								Type:        []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
						"queryStringWhitelist": {
							SchemaProps: spec.SchemaProps{
								Description: "Names of query string parameters to include in cache keys. All other parameters are excluded. Either specify QueryStringBlacklist or QueryStringWhitelist, but not both.",
								Type:        []string{"array"},
								Items: &spec.SchemaOrArray{
									Schema: &spec.Schema{
										SchemaProps: spec.SchemaProps{
											Type:   []string{"string"},
											Format: "",
										},
									},
								},
							},
						},
					},
				},
			},
			Dependencies: []string{},
		},
		"k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.IAPConfig": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "IAPConfig contains configuration for IAP-enabled backends.",
					Properties: map[string]spec.Schema{
						"enabled": {
							SchemaProps: spec.SchemaProps{
								Type:   []string{"boolean"},
								Format: "",
							},
						},
						"oauthclientCredentials": {
							SchemaProps: spec.SchemaProps{
								Ref: ref("k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.OAuthClientCredentials"),
							},
						},
					},
					Required: []string{"enabled", "oauthclientCredentials"},
				},
			},
			Dependencies: []string{
				"k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.OAuthClientCredentials"},
		},
		"k8s.io/ingress-gce/pkg/apis/backendconfig/v1beta1.OAuthClientCredentials": {
			Schema: spec.Schema{
				SchemaProps: spec.SchemaProps{
					Description: "OAuthClientCredentials contains credentials for a single IAP-enabled backend.",
					Properties: map[string]spec.Schema{
						"secretName": {
							SchemaProps: spec.SchemaProps{
								Description: "The name of a k8s secret which stores the OAuth client id & secret.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"clientID": {
							SchemaProps: spec.SchemaProps{
								Description: "Direct reference to OAuth client id.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
						"clientSecret": {
							SchemaProps: spec.SchemaProps{
								Description: "Direct reference to OAuth client secret.",
								Type:        []string{"string"},
								Format:      "",
							},
						},
					},
					Required: []string{"secretName"},
				},
			},
			Dependencies: []string{},
		},
	}
}
