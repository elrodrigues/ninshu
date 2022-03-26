package com

// To be implemented
import (
	"github.com/hashicorp/memberlist"
)

// options to be added
func Init() (*memberlist.Memberlist, error) {
	list, err := memberlist.Create(memberlist.DefaultLANConfig())
	return list, err
}
