package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listSubscriptionsCmd = &cobra.Command{
	Use:   "list-subscriptions",
	Short: "Lists all subscriptions created in an Amazon Q Business application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listSubscriptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_listSubscriptionsCmd).Standalone()

		qbusiness_listSubscriptionsCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application linked to the subscription.")
		qbusiness_listSubscriptionsCmd.Flags().String("max-results", "", "The maximum number of Amazon Q Business subscriptions to return.")
		qbusiness_listSubscriptionsCmd.Flags().String("next-token", "", "If the `maxResults` response was incomplete because there is more data to retrieve, Amazon Q Business returns a pagination token in the response.")
		qbusiness_listSubscriptionsCmd.MarkFlagRequired("application-id")
	})
	qbusinessCmd.AddCommand(qbusiness_listSubscriptionsCmd)
}
