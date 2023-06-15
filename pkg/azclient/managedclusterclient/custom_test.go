// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by client-gen. DO NOT EDIT.
package managedclusterclient

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"strings"

	"golang.org/x/crypto/ssh"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	armcontainerservice "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
	. "github.com/onsi/gomega"
)

func init() {
	addtionalTestCases = func() {
	}

	beforeAllFunc = func(ctx context.Context) {
		var generatedSSHKey string
		privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
		Expect(err).NotTo(HaveOccurred())

		// generate and write private key as PEM
		var privKeyBuf strings.Builder

		privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
		err = pem.Encode(&privKeyBuf, privateKeyPEM)
		Expect(err).NotTo(HaveOccurred())

		// generate and write public key
		pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
		Expect(err).NotTo(HaveOccurred())
		generatedSSHKey = string(ssh.MarshalAuthorizedKey(pub))
		newResource = &armcontainerservice.ManagedCluster{
			Location: to.Ptr(location),
			Identity: &armcontainerservice.ManagedClusterIdentity{
				Type: to.Ptr(armcontainerservice.ResourceIdentityTypeSystemAssigned),
			},
			SKU: &armcontainerservice.ManagedClusterSKU{
				Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierStandard),
				Name: to.Ptr(armcontainerservice.ManagedClusterSKUNameBase),
			},
			Properties: &armcontainerservice.ManagedClusterProperties{
				AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
					{
						Name:   to.Ptr("agentpool1"),
						Count:  to.Ptr[int32](3),
						VMSize: to.Ptr("Standard_DS2_v2"),
					},
				},
				DNSPrefix: to.Ptr(resourceGroupName + resourceName + "dnsPrefix"),
				LinuxProfile: &armcontainerservice.LinuxProfile{
					AdminUsername: to.Ptr("azureuser"),
					SSH: &armcontainerservice.SSHConfiguration{
						PublicKeys: []*armcontainerservice.SSHPublicKey{
							{
								KeyData: to.Ptr(generatedSSHKey),
							},
						},
					},
				},

				EnableRBAC: to.Ptr(true),
			},
		}
	}
	afterAllFunc = func(ctx context.Context) {
	}
}
