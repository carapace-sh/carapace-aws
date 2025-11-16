package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listFlowAssociationsCmd = &cobra.Command{
	Use:   "list-flow-associations",
	Short: "List the flow association based on the filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listFlowAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listFlowAssociationsCmd).Standalone()

		connect_listFlowAssociationsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listFlowAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listFlowAssociationsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listFlowAssociationsCmd.Flags().String("resource-type", "", "A valid resource type.")
		connect_listFlowAssociationsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listFlowAssociationsCmd)
}
