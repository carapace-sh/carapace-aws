package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getFlowAssociationCmd = &cobra.Command{
	Use:   "get-flow-association",
	Short: "Retrieves the flow associated for a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getFlowAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_getFlowAssociationCmd).Standalone()

		connect_getFlowAssociationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_getFlowAssociationCmd.Flags().String("resource-id", "", "The identifier of the resource.")
		connect_getFlowAssociationCmd.Flags().String("resource-type", "", "A valid resource type.")
		connect_getFlowAssociationCmd.MarkFlagRequired("instance-id")
		connect_getFlowAssociationCmd.MarkFlagRequired("resource-id")
		connect_getFlowAssociationCmd.MarkFlagRequired("resource-type")
	})
	connectCmd.AddCommand(connect_getFlowAssociationCmd)
}
