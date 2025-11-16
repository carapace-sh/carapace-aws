package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listContactFlowsCmd = &cobra.Command{
	Use:   "list-contact-flows",
	Short: "Provides information about the flows for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listContactFlowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listContactFlowsCmd).Standalone()

		connect_listContactFlowsCmd.Flags().String("contact-flow-types", "", "The type of flow.")
		connect_listContactFlowsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listContactFlowsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listContactFlowsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listContactFlowsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listContactFlowsCmd)
}
