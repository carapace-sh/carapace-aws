package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_dissociateEntityFromThingCmd = &cobra.Command{
	Use:   "dissociate-entity-from-thing",
	Short: "Dissociates a device entity from a concrete thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_dissociateEntityFromThingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotthingsgraph_dissociateEntityFromThingCmd).Standalone()

		iotthingsgraph_dissociateEntityFromThingCmd.Flags().String("entity-type", "", "The entity type from which to disassociate the thing.")
		iotthingsgraph_dissociateEntityFromThingCmd.Flags().String("thing-name", "", "The name of the thing to disassociate.")
		iotthingsgraph_dissociateEntityFromThingCmd.MarkFlagRequired("entity-type")
		iotthingsgraph_dissociateEntityFromThingCmd.MarkFlagRequired("thing-name")
	})
	iotthingsgraphCmd.AddCommand(iotthingsgraph_dissociateEntityFromThingCmd)
}
