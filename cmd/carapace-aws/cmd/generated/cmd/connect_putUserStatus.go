package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_putUserStatusCmd = &cobra.Command{
	Use:   "put-user-status",
	Short: "Changes the current status of a user or agent in Amazon Connect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_putUserStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_putUserStatusCmd).Standalone()

		connect_putUserStatusCmd.Flags().String("agent-status-id", "", "The identifier of the agent status.")
		connect_putUserStatusCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_putUserStatusCmd.Flags().String("user-id", "", "The identifier of the user.")
		connect_putUserStatusCmd.MarkFlagRequired("agent-status-id")
		connect_putUserStatusCmd.MarkFlagRequired("instance-id")
		connect_putUserStatusCmd.MarkFlagRequired("user-id")
	})
	connectCmd.AddCommand(connect_putUserStatusCmd)
}
