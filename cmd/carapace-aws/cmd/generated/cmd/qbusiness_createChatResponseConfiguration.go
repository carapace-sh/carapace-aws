package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_createChatResponseConfigurationCmd = &cobra.Command{
	Use:   "create-chat-response-configuration",
	Short: "Creates a new chat response configuration for an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_createChatResponseConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_createChatResponseConfigurationCmd).Standalone()

		qbusiness_createChatResponseConfigurationCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application for which to create the new chat response configuration.")
		qbusiness_createChatResponseConfigurationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
		qbusiness_createChatResponseConfigurationCmd.Flags().String("display-name", "", "A human-readable name for the new chat response configuration, making it easier to identify and manage among multiple configurations.")
		qbusiness_createChatResponseConfigurationCmd.Flags().String("response-configurations", "", "A collection of response configuration settings that define how Amazon Q Business will generate and format responses to user queries in chat interactions.")
		qbusiness_createChatResponseConfigurationCmd.Flags().String("tags", "", "A list of key-value pairs to apply as tags to the new chat response configuration, enabling categorization and management of resources across Amazon Web Services services.")
		qbusiness_createChatResponseConfigurationCmd.MarkFlagRequired("application-id")
		qbusiness_createChatResponseConfigurationCmd.MarkFlagRequired("display-name")
		qbusiness_createChatResponseConfigurationCmd.MarkFlagRequired("response-configurations")
	})
	qbusinessCmd.AddCommand(qbusiness_createChatResponseConfigurationCmd)
}
