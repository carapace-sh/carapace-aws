package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createAgentStatusCmd = &cobra.Command{
	Use:   "create-agent-status",
	Short: "Creates an agent status for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createAgentStatusCmd).Standalone()

	connect_createAgentStatusCmd.Flags().String("description", "", "The description of the status.")
	connect_createAgentStatusCmd.Flags().String("display-order", "", "The display order of the status.")
	connect_createAgentStatusCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createAgentStatusCmd.Flags().String("name", "", "The name of the status.")
	connect_createAgentStatusCmd.Flags().String("state", "", "The state of the status.")
	connect_createAgentStatusCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	connect_createAgentStatusCmd.MarkFlagRequired("instance-id")
	connect_createAgentStatusCmd.MarkFlagRequired("name")
	connect_createAgentStatusCmd.MarkFlagRequired("state")
	connectCmd.AddCommand(connect_createAgentStatusCmd)
}
