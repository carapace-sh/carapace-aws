package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_getChatControlsConfigurationCmd = &cobra.Command{
	Use:   "get-chat-controls-configuration",
	Short: "Gets information about chat controls configured for an existing Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_getChatControlsConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_getChatControlsConfigurationCmd).Standalone()

		qbusiness_getChatControlsConfigurationCmd.Flags().String("application-id", "", "The identifier of the application for which the chat controls are configured.")
		qbusiness_getChatControlsConfigurationCmd.Flags().String("max-results", "", "The maximum number of configured chat controls to return.")
		qbusiness_getChatControlsConfigurationCmd.Flags().String("next-token", "", "If the `maxResults` response was incomplete because there is more data to retrieve, Amazon Q Business returns a pagination token in the response.")
		qbusiness_getChatControlsConfigurationCmd.MarkFlagRequired("application-id")
	})
	qbusinessCmd.AddCommand(qbusiness_getChatControlsConfigurationCmd)
}
