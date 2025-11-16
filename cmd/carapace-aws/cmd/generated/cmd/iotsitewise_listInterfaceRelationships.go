package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listInterfaceRelationshipsCmd = &cobra.Command{
	Use:   "list-interface-relationships",
	Short: "Retrieves a paginated list of asset models that have a specific interface asset model applied to them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listInterfaceRelationshipsCmd).Standalone()

	iotsitewise_listInterfaceRelationshipsCmd.Flags().String("interface-asset-model-id", "", "The ID of the interface asset model.")
	iotsitewise_listInterfaceRelationshipsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
	iotsitewise_listInterfaceRelationshipsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	iotsitewise_listInterfaceRelationshipsCmd.MarkFlagRequired("interface-asset-model-id")
	iotsitewiseCmd.AddCommand(iotsitewise_listInterfaceRelationshipsCmd)
}
