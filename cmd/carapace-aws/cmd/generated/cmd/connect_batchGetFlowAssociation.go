package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_batchGetFlowAssociationCmd = &cobra.Command{
	Use:   "batch-get-flow-association",
	Short: "Retrieve the flow associations for the given resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_batchGetFlowAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_batchGetFlowAssociationCmd).Standalone()

		connect_batchGetFlowAssociationCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_batchGetFlowAssociationCmd.Flags().String("resource-ids", "", "A list of resource identifiers to retrieve flow associations.")
		connect_batchGetFlowAssociationCmd.Flags().String("resource-type", "", "The type of resource association.")
		connect_batchGetFlowAssociationCmd.MarkFlagRequired("instance-id")
		connect_batchGetFlowAssociationCmd.MarkFlagRequired("resource-ids")
	})
	connectCmd.AddCommand(connect_batchGetFlowAssociationCmd)
}
