package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_deleteEventSubscriptionCmd = &cobra.Command{
	Use:   "delete-event-subscription",
	Short: "Deletes an Amazon DocumentDB event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_deleteEventSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_deleteEventSubscriptionCmd).Standalone()

		docdb_deleteEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the Amazon DocumentDB event notification subscription that you want to delete.")
		docdb_deleteEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	})
	docdbCmd.AddCommand(docdb_deleteEventSubscriptionCmd)
}
