package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listAgentStatusesCmd = &cobra.Command{
	Use:   "list-agent-statuses",
	Short: "Lists agent statuses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listAgentStatusesCmd).Standalone()

	connect_listAgentStatusesCmd.Flags().String("agent-status-types", "", "Available agent status types.")
	connect_listAgentStatusesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_listAgentStatusesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_listAgentStatusesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_listAgentStatusesCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_listAgentStatusesCmd)
}
