package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatbot_disassociateFromConfigurationCmd = &cobra.Command{
	Use:   "disassociate-from-configuration",
	Short: "Unlink a resource, for example a custom action, from a channel configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatbot_disassociateFromConfigurationCmd).Standalone()

	chatbot_disassociateFromConfigurationCmd.Flags().String("chat-configuration", "", "The channel configuration the resource is being disassociated from.")
	chatbot_disassociateFromConfigurationCmd.Flags().String("resource", "", "The resource (for example, a custom action) Amazon Resource Name (ARN) to unlink.")
	chatbot_disassociateFromConfigurationCmd.MarkFlagRequired("chat-configuration")
	chatbot_disassociateFromConfigurationCmd.MarkFlagRequired("resource")
	chatbotCmd.AddCommand(chatbot_disassociateFromConfigurationCmd)
}
