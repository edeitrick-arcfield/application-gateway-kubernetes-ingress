// -------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// --------------------------------------------------------------------------------------------

package appgw

import (
	"context"
	"fmt"
	"strings"
	"time"

	n "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2021-03-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	networking "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	testclient "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/record"

	"github.com/Azure/application-gateway-kubernetes-ingress/pkg/annotations"
	"github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/agic_crd_client/clientset/versioned/fake"
	multiCluster_fake "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/azure_multicluster_crd_client/clientset/versioned/fake"
	istio_fake "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/istio_crd_client/clientset/versioned/fake"
	"github.com/Azure/application-gateway-kubernetes-ingress/pkg/environment"
	"github.com/Azure/application-gateway-kubernetes-ingress/pkg/k8scontext"
	"github.com/Azure/application-gateway-kubernetes-ingress/pkg/metricstore"
	"github.com/Azure/application-gateway-kubernetes-ingress/pkg/tests"
	"github.com/Azure/application-gateway-kubernetes-ingress/pkg/tests/mocks"
	"github.com/Azure/application-gateway-kubernetes-ingress/pkg/utils"
	"github.com/Azure/application-gateway-kubernetes-ingress/pkg/version"
)

var _ = Describe("Public and Private IP tests", func() {
	var configBuilder ConfigBuilder
	var appGwy *n.ApplicationGateway
	var cbCtx *ConfigBuilderContext
	version.Version = "a"
	version.GitCommit = "b"
	version.BuildDate = "c"

	ingressNS := tests.Namespace

	// Create the "test-ingressPrivateIP-controller" namespace.
	// We will create all our resources under this namespace.
	nameSpace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: ingressNS,
		},
	}

	// Create a node
	node := &v1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "node-1",
		},
		Spec: v1.NodeSpec{
			ProviderID: "azure:///subscriptions/subid/resourceGroups/MC_aksresgp_aksname_location/providers/Microsoft.Compute/virtualMachines/vmname",
		},
	}

	ingressPublicIP := &networking.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotations.IngressClassKey: environment.DefaultIngressClassController,
			},
			Namespace: ingressNS,
			Name:      "external-ingress-resource",
		},
		Spec: networking.IngressSpec{
			Rules: []networking.IngressRule{
				{
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{
								{
									Path: "/*",
									Backend: networking.IngressBackend{
										Service: &networking.IngressServiceBackend{
											Name: tests.ServiceName,
											Port: networking.ServiceBackendPort{
												Number: 443,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			TLS: []networking.IngressTLS{
				{
					Hosts: []string{
						"pub.lic",
						"www.contoso.com",
						"ftp.contoso.com",
						tests.Host,
						"",
					},
					SecretName: tests.NameOfSecret,
				},
				{
					Hosts:      []string{},
					SecretName: tests.NameOfSecret,
				},
			},
		},
	}

	// Create the Ingress resource.
	ingressPrivateIP := &networking.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				annotations.IngressClassKey: environment.DefaultIngressClassController,
				annotations.UsePrivateIPKey: "true",
			},
			Namespace: ingressNS,
			Name:      "internal-ingress-resource",
		},
		Spec: networking.IngressSpec{
			Rules: []networking.IngressRule{
				{
					IngressRuleValue: networking.IngressRuleValue{
						HTTP: &networking.HTTPIngressRuleValue{
							Paths: []networking.HTTPIngressPath{
								{
									Path: "/*",
									Backend: networking.IngressBackend{
										Service: &networking.IngressServiceBackend{
											Name: tests.ServiceName,
											Port: networking.ServiceBackendPort{
												Number: 80,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tests.ServiceName,
			Namespace: ingressNS,
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Name: "http",
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 80,
					},
					Protocol: v1.ProtocolTCP,
					Port:     int32(80),
				},
				{
					Name: "https",
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 80,
					},
					Protocol: v1.ProtocolTCP,
					Port:     int32(443),
				},
			},
			Selector: map[string]string{"app": "web--app--name"},
		},
	}

	serviceList := []*v1.Service{
		service,
	}

	endpoints := &v1.Endpoints{
		ObjectMeta: metav1.ObjectMeta{
			Name:      tests.ServiceName,
			Namespace: ingressNS,
		},
		Subsets: []v1.EndpointSubset{
			{
				Addresses: []v1.EndpointAddress{
					{IP: "1.1.1.1"},
					{IP: "1.1.1.2"},
					{IP: "1.1.1.3"},
				},
				Ports: []v1.EndpointPort{
					{
						Name:     "http",
						Port:     int32(80),
						Protocol: v1.ProtocolTCP,
					},
					{
						Name:     "https",
						Port:     int32(443),
						Protocol: v1.ProtocolTCP,
					},
				},
			},
		},
	}

	pod1 := tests.NewPodFixture("pod1", ingressNS, "http", int32(80))
	pod2 := tests.NewPodFixture("pod2", ingressNS, "https", int32(80))

	appGwIdentifier := Identifier{
		SubscriptionID: tests.Subscription,
		ResourceGroup:  tests.ResourceGroup,
		AppGwName:      tests.AppGwName,
	}

	// Create the mock K8s client.
	k8sClient := testclient.NewSimpleClientset()

	It("should have not failed", func() {
		_, err := k8sClient.CoreV1().Namespaces().Create(context.TODO(), nameSpace, metav1.CreateOptions{})
		Expect(err).ToNot(HaveOccurred())
	})

	crdClient := fake.NewSimpleClientset()
	istioCrdClient := istio_fake.NewSimpleClientset()
	multiClusterCrdClient := multiCluster_fake.NewSimpleClientset()
	// Removed the reference to IsNetworkingV1PackageSupported as K8 versions prior to V1.19 are no longer supported V1.19 has not been supported itself
	// since around August 2021 we should not be supporting overley deprecated version that are out of support
	ctxt := k8scontext.NewContext(k8sClient, crdClient, multiClusterCrdClient, istioCrdClient, []string{ingressNS}, 1000*time.Second, metricstore.NewFakeMetricStore(), environment.GetFakeEnv())

	secret := tests.NewSecretTestFixture()

	err := ctxt.Caches.Secret.Add(secret)
	It("should have not failed", func() { Expect(err).ToNot(HaveOccurred()) })

	secKey := utils.GetResourceKey(secret.Namespace, secret.Name)

	err = ctxt.CertificateSecretStore.ConvertSecret(secKey, secret)
	It("should have converted the certificate", func() { Expect(err).ToNot(HaveOccurred()) })

	pfx := ctxt.CertificateSecretStore.GetPfxCertificate(secKey)
	It("should have found the pfx certificate", func() { Expect(pfx).ToNot(BeNil()) })

	ctxtSecret := ctxt.GetSecret(secKey)
	It("should have found the secret", func() { Expect(ctxtSecret).To(Equal(secret)) })

	_, err = k8sClient.CoreV1().Nodes().Create(context.TODO(), node, metav1.CreateOptions{})
	It("should have not failed", func() { Expect(err).ToNot(HaveOccurred()) })

	_, err = k8sClient.NetworkingV1().Ingresses(ingressNS).Create(context.TODO(), ingressPublicIP, metav1.CreateOptions{})
	It("should have not failed", func() { Expect(err).ToNot(HaveOccurred()) })

	_, err = k8sClient.NetworkingV1().Ingresses(ingressNS).Update(context.TODO(), ingressPublicIP, metav1.UpdateOptions{})
	It("should have not failed", func() { Expect(err).ToNot(HaveOccurred(), "Unable to update ingress resource due to: %v", err) })

	_, err = k8sClient.NetworkingV1().Ingresses(ingressNS).Create(context.TODO(), ingressPrivateIP, metav1.CreateOptions{})
	It("should have not failed", func() { Expect(err).ToNot(HaveOccurred()) })

	_, err = k8sClient.NetworkingV1().Ingresses(ingressNS).Update(context.TODO(), ingressPrivateIP, metav1.UpdateOptions{})
	It("should have not failed", func() { Expect(err).ToNot(HaveOccurred(), "Unable to update ingress resource due to: %v", err) })

	_, err = k8sClient.CoreV1().Services(ingressNS).Create(context.TODO(), service, metav1.CreateOptions{})
	It("should have not failed", func() { Expect(err).ToNot(HaveOccurred()) })

	_, err = k8sClient.CoreV1().Endpoints(ingressNS).Create(context.TODO(), endpoints, metav1.CreateOptions{})
	It("should have not failed", func() { Expect(err).ToNot(HaveOccurred()) })

	_, err = k8sClient.CoreV1().Pods(ingressNS).Create(context.TODO(), pod1, metav1.CreateOptions{})
	It("should have not failed", func() { Expect(err).ToNot(HaveOccurred()) })

	_, err = k8sClient.CoreV1().Pods(ingressNS).Create(context.TODO(), pod2, metav1.CreateOptions{})
	It("should have not failed", func() { Expect(err).ToNot(HaveOccurred()) })

	Context("Both private ip and public ip are present on the Gateway", func() {
		BeforeEach(func() {
			appGwy = &n.ApplicationGateway{
				ApplicationGatewayPropertiesFormat: NewAppGwyConfigFixture(),
			}

			appGwy.FrontendIPConfigurations = &[]n.ApplicationGatewayFrontendIPConfiguration{
				{
					// Public IP
					Name: to.StringPtr("public"),
					ID:   to.StringPtr("public"),
					ApplicationGatewayFrontendIPConfigurationPropertiesFormat: &n.ApplicationGatewayFrontendIPConfigurationPropertiesFormat{
						PrivateIPAddress: nil,
						PublicIPAddress: &n.SubResource{
							ID: to.StringPtr("xyz"),
						},
					},
				},
				{
					// Private IP
					Name: to.StringPtr("private"),
					ID:   to.StringPtr("private"),
					ApplicationGatewayFrontendIPConfigurationPropertiesFormat: &n.ApplicationGatewayFrontendIPConfigurationPropertiesFormat{
						PrivateIPAddress: to.StringPtr("abc"),
						PublicIPAddress:  nil,
					},
				},
			}

			configBuilder = NewConfigBuilder(ctxt, &appGwIdentifier, appGwy, record.NewFakeRecorder(100), mocks.Clock{})
			cbCtx = &ConfigBuilderContext{
				IngressList: []*networking.Ingress{
					ingressPrivateIP,
					ingressPublicIP,
				},
				ServiceList:           serviceList,
				EnvVariables:          environment.GetFakeEnv(),
				DefaultAddressPoolID:  to.StringPtr("xx"),
				DefaultHTTPSettingsID: to.StringPtr("yy"),
			}
		})

		It("should use private ip for listener at port 80 as ingress is using use-private-ip annotation", func() {
			appGW, err := configBuilder.Build(cbCtx)
			Expect(err).ToNot(HaveOccurred())

			jsonBlob, err := appGW.MarshalJSON()
			Expect(err).ToNot(HaveOccurred())

			Expect(appGW.HTTPListeners).ToNot(BeNil())

			foundPrivateIPListener := false
			foundPublicIPListener := false
			for _, listener := range *appGW.HTTPListeners {
				// port 80 should be used with private ip
				if strings.Contains(*listener.FrontendPort.ID, "fp-80") {
					foundPrivateIPListener = true
					Expect(*listener.FrontendIPConfiguration.ID).To(Equal("private"), "expecting to find private IP frontend configuration attached here")
				}

				// port 443 should be used with public ip
				if strings.Contains(*listener.FrontendPort.ID, "fp-443") {
					foundPublicIPListener = true
					Expect(*listener.FrontendIPConfiguration.ID).To(Equal("public"), "expecting to find public IP frontend configuration attached here")
				}
			}

			Expect(foundPrivateIPListener).To(BeTrue(), fmt.Sprintf("Expecting to find a listener using private IP. Actual JSON:\n%s\n", string(jsonBlob)))
			Expect(foundPublicIPListener).To(BeTrue(), fmt.Sprintf("Expecting to find a listener using private IP. Actual JSON:\n%s\n", string(jsonBlob)))
		})
	})

	Context("Only private ip present on the Gateway", func() {
		BeforeEach(func() {
			appGwy = &n.ApplicationGateway{
				ApplicationGatewayPropertiesFormat: NewAppGwyConfigFixture(),
			}

			appGwy.FrontendIPConfigurations = &[]n.ApplicationGatewayFrontendIPConfiguration{
				{
					// Private IP
					Name: to.StringPtr("private"),
					ID:   to.StringPtr("private"),
					ApplicationGatewayFrontendIPConfigurationPropertiesFormat: &n.ApplicationGatewayFrontendIPConfigurationPropertiesFormat{
						PrivateIPAddress: to.StringPtr("abc"),
						PublicIPAddress:  nil,
					},
				},
			}

			configBuilder = NewConfigBuilder(ctxt, &appGwIdentifier, appGwy, record.NewFakeRecorder(100), mocks.Clock{})
			cbCtx = &ConfigBuilderContext{
				IngressList: []*networking.Ingress{
					ingressPrivateIP,
				},
				ServiceList:           serviceList,
				EnvVariables:          environment.GetFakeEnv(),
				DefaultAddressPoolID:  to.StringPtr("xx"),
				DefaultHTTPSettingsID: to.StringPtr("yy"),
			}
		})

		It("should use private ip for all listeners and for all ingresses", func() {
			appGW, err := configBuilder.Build(cbCtx)
			Expect(err).ToNot(HaveOccurred())

			Expect(appGW.HTTPListeners).ToNot(BeNil())
			Expect(len(*appGW.HTTPListeners)).To(Equal(1))

			for _, listener := range *appGW.HTTPListeners {
				Expect(*listener.FrontendIPConfiguration.ID).To(Equal("private"), "expecting to find private IP frontend configuration attached here")
			}

			Expect(appGW.RequestRoutingRules).ToNot(BeNil())
			Expect(len(*appGW.RequestRoutingRules)).To(Equal(1), "should have 2 rules")
		})
	})
})
