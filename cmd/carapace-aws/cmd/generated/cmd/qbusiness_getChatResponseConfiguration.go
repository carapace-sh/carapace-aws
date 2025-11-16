package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getChatResponseConfigurationCmd = &cobra.Command{
	Use:   "get-chat-response-configuration",
	Short: "Retrieves detailed information about a specific chat response configuration from an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getChatResponseConfigurationCmd).Standalone()

	qbusiness_getChatResponseConfigurationCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application containing the chat response configuration to retrieve.")
	qbusiness_getChatResponseConfigurationCmd.Flags().String("chat-response-configuration-id", "", "The unique identifier of the chat response configuration to retrieve from the specified application.")
	qbusiness_getChatResponseConfigurationCmd.MarkFlagRequired("application-id")
	qbusiness_getChatResponseConfigurationCmd.MarkFlagRequired("chat-response-configuration-id")
	qbusinessCmd.AddCommand(qbusiness_getChatResponseConfigurationCmd)
}
