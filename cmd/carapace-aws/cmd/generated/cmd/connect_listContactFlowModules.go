package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listContactFlowModulesCmd = &cobra.Command{
	Use:   "list-contact-flow-modules",
	Short: "Provides information about the flow modules for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listContactFlowModulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listContactFlowModulesCmd).Standalone()

		connect_listContactFlowModulesCmd.Flags().String("contact-flow-module-state", "", "The state of the flow module.")
		connect_listContactFlowModulesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listContactFlowModulesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listContactFlowModulesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listContactFlowModulesCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listContactFlowModulesCmd)
}
