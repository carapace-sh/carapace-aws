package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listConversationsCmd = &cobra.Command{
	Use:   "list-conversations",
	Short: "Lists one or more Amazon Q Business conversations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listConversationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_listConversationsCmd).Standalone()

		qbusiness_listConversationsCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application.")
		qbusiness_listConversationsCmd.Flags().String("max-results", "", "The maximum number of Amazon Q Business conversations to return.")
		qbusiness_listConversationsCmd.Flags().String("next-token", "", "If the `maxResults` response was incomplete because there is more data to retrieve, Amazon Q Business returns a pagination token in the response.")
		qbusiness_listConversationsCmd.Flags().String("user-id", "", "The identifier of the user involved in the Amazon Q Business web experience conversation.")
		qbusiness_listConversationsCmd.MarkFlagRequired("application-id")
	})
	qbusinessCmd.AddCommand(qbusiness_listConversationsCmd)
}
