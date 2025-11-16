package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_updateChatControlsConfigurationCmd = &cobra.Command{
	Use:   "update-chat-controls-configuration",
	Short: "Updates a set of chat controls configured for an existing Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_updateChatControlsConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_updateChatControlsConfigurationCmd).Standalone()

		qbusiness_updateChatControlsConfigurationCmd.Flags().String("application-id", "", "The identifier of the application for which the chat controls are configured.")
		qbusiness_updateChatControlsConfigurationCmd.Flags().String("blocked-phrases-configuration-update", "", "The phrases blocked from chat by your chat control configuration.")
		qbusiness_updateChatControlsConfigurationCmd.Flags().String("client-token", "", "A token that you provide to identify the request to update a Amazon Q Business application chat configuration.")
		qbusiness_updateChatControlsConfigurationCmd.Flags().String("creator-mode-configuration", "", "The configuration details for `CREATOR_MODE`.")
		qbusiness_updateChatControlsConfigurationCmd.Flags().String("hallucination-reduction-configuration", "", "The hallucination reduction settings for your application.")
		qbusiness_updateChatControlsConfigurationCmd.Flags().String("orchestration-configuration", "", "The chat response orchestration settings for your application.")
		qbusiness_updateChatControlsConfigurationCmd.Flags().String("response-scope", "", "The response scope configured for your application.")
		qbusiness_updateChatControlsConfigurationCmd.Flags().String("topic-configurations-to-create-or-update", "", "The configured topic specific chat controls you want to update.")
		qbusiness_updateChatControlsConfigurationCmd.Flags().String("topic-configurations-to-delete", "", "The configured topic specific chat controls you want to delete.")
		qbusiness_updateChatControlsConfigurationCmd.MarkFlagRequired("application-id")
	})
	qbusinessCmd.AddCommand(qbusiness_updateChatControlsConfigurationCmd)
}
