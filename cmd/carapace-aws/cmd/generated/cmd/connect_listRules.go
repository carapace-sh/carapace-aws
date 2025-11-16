package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listRulesCmd = &cobra.Command{
	Use:   "list-rules",
	Short: "List all rules for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listRulesCmd).Standalone()

		connect_listRulesCmd.Flags().String("event-source-name", "", "The name of the event source.")
		connect_listRulesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listRulesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listRulesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listRulesCmd.Flags().String("publish-status", "", "The publish status of the rule.")
		connect_listRulesCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listRulesCmd)
}
