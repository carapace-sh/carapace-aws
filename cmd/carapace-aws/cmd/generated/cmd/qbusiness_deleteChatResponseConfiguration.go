package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_deleteChatResponseConfigurationCmd = &cobra.Command{
	Use:   "delete-chat-response-configuration",
	Short: "Deletes a specified chat response configuration from an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_deleteChatResponseConfigurationCmd).Standalone()

	qbusiness_deleteChatResponseConfigurationCmd.Flags().String("application-id", "", "The unique identifier of theAmazon Q Business application from which to delete the chat response configuration.")
	qbusiness_deleteChatResponseConfigurationCmd.Flags().String("chat-response-configuration-id", "", "The unique identifier of the chat response configuration to delete from the specified application.")
	qbusiness_deleteChatResponseConfigurationCmd.MarkFlagRequired("application-id")
	qbusiness_deleteChatResponseConfigurationCmd.MarkFlagRequired("chat-response-configuration-id")
	qbusinessCmd.AddCommand(qbusiness_deleteChatResponseConfigurationCmd)
}
