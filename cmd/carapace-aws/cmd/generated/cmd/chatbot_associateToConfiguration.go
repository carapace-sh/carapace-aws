package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_associateToConfigurationCmd = &cobra.Command{
	Use:   "associate-to-configuration",
	Short: "Links a resource (for example, a custom action) to a channel configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_associateToConfigurationCmd).Standalone()

	chatbot_associateToConfigurationCmd.Flags().String("chat-configuration", "", "The channel configuration to associate with the resource.")
	chatbot_associateToConfigurationCmd.Flags().String("resource", "", "The resource Amazon Resource Name (ARN) to link.")
	chatbot_associateToConfigurationCmd.MarkFlagRequired("chat-configuration")
	chatbot_associateToConfigurationCmd.MarkFlagRequired("resource")
	chatbotCmd.AddCommand(chatbot_associateToConfigurationCmd)
}
