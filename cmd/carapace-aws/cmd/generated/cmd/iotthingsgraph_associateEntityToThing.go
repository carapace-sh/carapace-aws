package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_associateEntityToThingCmd = &cobra.Command{
	Use:   "associate-entity-to-thing",
	Short: "Associates a device with a concrete thing that is in the user's registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_associateEntityToThingCmd).Standalone()

	iotthingsgraph_associateEntityToThingCmd.Flags().String("entity-id", "", "The ID of the device to be associated with the thing.")
	iotthingsgraph_associateEntityToThingCmd.Flags().String("namespace-version", "", "The version of the user's namespace.")
	iotthingsgraph_associateEntityToThingCmd.Flags().String("thing-name", "", "The name of the thing to which the entity is to be associated.")
	iotthingsgraph_associateEntityToThingCmd.MarkFlagRequired("entity-id")
	iotthingsgraph_associateEntityToThingCmd.MarkFlagRequired("thing-name")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_associateEntityToThingCmd)
}
