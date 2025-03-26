// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	projectcalicov3 "github.com/projectcalico/api/pkg/client/clientset_generated/clientset/typed/projectcalico/v3"
	gentype "k8s.io/client-go/gentype"
)

// fakeBGPConfigurations implements BGPConfigurationInterface
type fakeBGPConfigurations struct {
	*gentype.FakeClientWithList[*v3.BGPConfiguration, *v3.BGPConfigurationList]
	Fake *FakeProjectcalicoV3
}

func newFakeBGPConfigurations(fake *FakeProjectcalicoV3) projectcalicov3.BGPConfigurationInterface {
	return &fakeBGPConfigurations{
		gentype.NewFakeClientWithList[*v3.BGPConfiguration, *v3.BGPConfigurationList](
			fake.Fake,
			"",
			v3.SchemeGroupVersion.WithResource("bgpconfigurations"),
			v3.SchemeGroupVersion.WithKind("BGPConfiguration"),
			func() *v3.BGPConfiguration { return &v3.BGPConfiguration{} },
			func() *v3.BGPConfigurationList { return &v3.BGPConfigurationList{} },
			func(dst, src *v3.BGPConfigurationList) { dst.ListMeta = src.ListMeta },
			func(list *v3.BGPConfigurationList) []*v3.BGPConfiguration { return gentype.ToPointerSlice(list.Items) },
			func(list *v3.BGPConfigurationList, items []*v3.BGPConfiguration) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
