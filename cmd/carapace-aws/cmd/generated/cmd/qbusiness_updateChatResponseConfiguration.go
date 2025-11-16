package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_updateChatResponseConfigurationCmd = &cobra.Command{
	Use:   "update-chat-response-configuration",
	Short: "Updates an existing chat response configuration in an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_updateChatResponseConfigurationCmd).Standalone()

	qbusiness_updateChatResponseConfigurationCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application containing the chat response configuration to update.")
	qbusiness_updateChatResponseConfigurationCmd.Flags().String("chat-response-configuration-id", "", "The unique identifier of the chat response configuration to update within the specified application.")
	qbusiness_updateChatResponseConfigurationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
	qbusiness_updateChatResponseConfigurationCmd.Flags().String("display-name", "", "The new human-readable name to assign to the chat response configuration, making it easier to identify among multiple configurations.")
	qbusiness_updateChatResponseConfigurationCmd.Flags().String("response-configurations", "", "The updated collection of response configuration settings that define how Amazon Q Business generates and formats responses to user queries.")
	qbusiness_updateChatResponseConfigurationCmd.MarkFlagRequired("application-id")
	qbusiness_updateChatResponseConfigurationCmd.MarkFlagRequired("chat-response-configuration-id")
	qbusiness_updateChatResponseConfigurationCmd.MarkFlagRequired("response-configurations")
	qbusinessCmd.AddCommand(qbusiness_updateChatResponseConfigurationCmd)
}
