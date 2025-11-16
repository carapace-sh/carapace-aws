package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deleteChatControlsConfigurationCmd = &cobra.Command{
	Use:   "delete-chat-controls-configuration",
	Short: "Deletes chat controls configured for an existing Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deleteChatControlsConfigurationCmd).Standalone()

	qbusiness_deleteChatControlsConfigurationCmd.Flags().String("application-id", "", "The identifier of the application the chat controls have been configured for.")
	qbusiness_deleteChatControlsConfigurationCmd.MarkFlagRequired("application-id")
	qbusinessCmd.AddCommand(qbusiness_deleteChatControlsConfigurationCmd)
}
