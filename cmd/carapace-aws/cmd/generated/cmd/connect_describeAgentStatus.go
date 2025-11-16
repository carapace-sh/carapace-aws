package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeAgentStatusCmd = &cobra.Command{
	Use:   "describe-agent-status",
	Short: "Describes an agent status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeAgentStatusCmd).Standalone()

	connect_describeAgentStatusCmd.Flags().String("agent-status-id", "", "The identifier for the agent status.")
	connect_describeAgentStatusCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeAgentStatusCmd.MarkFlagRequired("agent-status-id")
	connect_describeAgentStatusCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_describeAgentStatusCmd)
}
