// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
)

// IPReservationLister helps list IPReservations.
// All objects returned here must be treated as read-only.
type IPReservationLister interface {
	// List lists all IPReservations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.IPReservation, err error)
	// Get retrieves the IPReservation from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.IPReservation, error)
	IPReservationListerExpansion
}

// iPReservationLister implements the IPReservationLister interface.
type iPReservationLister struct {
	indexer cache.Indexer
}

// NewIPReservationLister returns a new IPReservationLister.
func NewIPReservationLister(indexer cache.Indexer) IPReservationLister {
	return &iPReservationLister{indexer: indexer}
}

// List lists all IPReservations in the indexer.
func (s *iPReservationLister) List(selector labels.Selector) (ret []*v3.IPReservation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.IPReservation))
	})
	return ret, err
}

// Get retrieves the IPReservation from the index for a given name.
func (s *iPReservationLister) Get(name string) (*v3.IPReservation, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("ipreservation"), name)
	}
	return obj.(*v3.IPReservation), nil
}