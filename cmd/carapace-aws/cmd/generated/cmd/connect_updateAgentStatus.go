package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateAgentStatusCmd = &cobra.Command{
	Use:   "update-agent-status",
	Short: "Updates agent status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateAgentStatusCmd).Standalone()

	connect_updateAgentStatusCmd.Flags().String("agent-status-id", "", "The identifier of the agent status.")
	connect_updateAgentStatusCmd.Flags().String("description", "", "The description of the agent status.")
	connect_updateAgentStatusCmd.Flags().String("display-order", "", "The display order of the agent status.")
	connect_updateAgentStatusCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_updateAgentStatusCmd.Flags().String("name", "", "The name of the agent status.")
	connect_updateAgentStatusCmd.Flags().Bool("no-reset-order-number", false, "A number indicating the reset order of the agent status.")
	connect_updateAgentStatusCmd.Flags().Bool("reset-order-number", false, "A number indicating the reset order of the agent status.")
	connect_updateAgentStatusCmd.Flags().String("state", "", "The state of the agent status.")
	connect_updateAgentStatusCmd.MarkFlagRequired("agent-status-id")
	connect_updateAgentStatusCmd.MarkFlagRequired("instance-id")
	connect_updateAgentStatusCmd.Flag("no-reset-order-number").Hidden = true
	connectCmd.AddCommand(connect_updateAgentStatusCmd)
}
