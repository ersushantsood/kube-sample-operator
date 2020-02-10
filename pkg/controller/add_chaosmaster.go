package controller

import (
	"github.com/ersushantsood/kube-sample-operator/pkg/controller/chaosmaster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, chaosmaster.Add)
}
