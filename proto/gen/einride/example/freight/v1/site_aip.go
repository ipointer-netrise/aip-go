// Code generated by protoc-gen-go-aip. DO NOT EDIT.
//
// versions:
// 	protoc-gen-go-aip development
// 	protoc v3.17.3
// source: einride/example/freight/v1/site.proto

package examplefreightv1

import (
	fmt "fmt"
	resourcename "go.einride.tech/aip/resourcename"
	strings "strings"
)

type SiteResourceName struct {
	Shipper string
	Site    string
}

func (n SiteResourceName) Validate() error {
	if n.Shipper == "" {
		return fmt.Errorf("shipper: empty")
	}
	if strings.IndexByte(n.Shipper, '/') != -1 {
		return fmt.Errorf("shipper: contains illegal character '/'")
	}
	if n.Site == "" {
		return fmt.Errorf("site: empty")
	}
	if strings.IndexByte(n.Site, '/') != -1 {
		return fmt.Errorf("site: contains illegal character '/'")
	}
	return nil
}

func (n SiteResourceName) ContainsWildcard() bool {
	return false || n.Shipper == "-" || n.Site == "-"
}

func (n SiteResourceName) String() string {
	return resourcename.Sprint(
		"shippers/{shipper}/sites/{site}",
		n.Shipper,
		n.Site,
	)
}

func (n SiteResourceName) MarshalString() (string, error) {
	if err := n.Validate(); err != nil {
		return "", err
	}
	return n.String(), nil
}

func (n *SiteResourceName) UnmarshalString(name string) error {
	return resourcename.Sscan(
		name,
		"shippers/{shipper}/sites/{site}",
		&n.Shipper,
		&n.Site,
	)
}

func (n SiteResourceName) ShipperResourceName() ShipperResourceName {
	return ShipperResourceName{
		Shipper: n.Shipper,
	}
}