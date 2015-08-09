package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/mikerjacobi/searches/models"
)

func main() {
	root := models.NewBNode(5)
	t := models.NewBTree(root)

	n1 := models.NewBNode(1).(*models.BNode)
	err := t.AddNode(n1)
	if err != nil {
		fmt.Sprintf(err.Error())
		panic(err)
	}

	found, _, err := models.BSearch(t, 1)
	if err != nil {
		fmt.Sprintf(err.Error())
		panic(err)
	}

	logrus.Infof("found: %t", found)
}
