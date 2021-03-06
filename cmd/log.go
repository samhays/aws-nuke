package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rebuy-de/aws-nuke/resources"
)

var (
	ReasonSkip            = *color.New(color.FgYellow)
	ReasonError           = *color.New(color.FgRed)
	ReasonRemoveTriggered = *color.New(color.FgGreen)
	ReasonWaitPending     = *color.New(color.FgBlue)
	ReasonSuccess         = *color.New(color.FgGreen)
)

var (
	ColorRegion       = *color.New(color.Bold)
	ColorResourceType = *color.New()
	ColorResourceID   = *color.New(color.Bold)
)

func Log(region Region, resourceType string, r resources.Resource, c color.Color, msg string) {
	ColorRegion.Printf("%s", region.Name)
	fmt.Printf(" - ")
	ColorResourceType.Print(resourceType)
	fmt.Printf(" - ")
	ColorResourceID.Printf("'%s'", r.String())
	fmt.Printf(" - ")
	c.Printf("%s\n", msg)
}
