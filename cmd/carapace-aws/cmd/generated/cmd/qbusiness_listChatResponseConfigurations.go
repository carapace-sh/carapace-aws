package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listChatResponseConfigurationsCmd = &cobra.Command{
	Use:   "list-chat-response-configurations",
	Short: "Retrieves a list of all chat response configurations available in a specified Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listChatResponseConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_listChatResponseConfigurationsCmd).Standalone()

		qbusiness_listChatResponseConfigurationsCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application for which to list available chat response configurations.")
		qbusiness_listChatResponseConfigurationsCmd.Flags().String("max-results", "", "The maximum number of chat response configurations to return in a single response.")
		qbusiness_listChatResponseConfigurationsCmd.Flags().String("next-token", "", "A pagination token used to retrieve the next set of results when the number of configurations exceeds the specified `maxResults` value.")
		qbusiness_listChatResponseConfigurationsCmd.MarkFlagRequired("application-id")
	})
	qbusinessCmd.AddCommand(qbusiness_listChatResponseConfigurationsCmd)
}
