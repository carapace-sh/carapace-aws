package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listContactFlowVersionsCmd = &cobra.Command{
	Use:   "list-contact-flow-versions",
	Short: "Returns all the available versions for the specified Amazon Connect instance and flow identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listContactFlowVersionsCmd).Standalone()

	connect_listContactFlowVersionsCmd.Flags().String("contact-flow-id", "", "The identifier of the flow.")
	connect_listContactFlowVersionsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listContactFlowVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listContactFlowVersionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listContactFlowVersionsCmd.MarkFlagRequired("contact-flow-id")
	connect_listContactFlowVersionsCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_listContactFlowVersionsCmd)
}
