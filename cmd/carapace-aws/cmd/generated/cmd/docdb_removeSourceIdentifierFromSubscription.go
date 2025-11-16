package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_removeSourceIdentifierFromSubscriptionCmd = &cobra.Command{
	Use:   "remove-source-identifier-from-subscription",
	Short: "Removes a source identifier from an existing Amazon DocumentDB event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_removeSourceIdentifierFromSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_removeSourceIdentifierFromSubscriptionCmd).Standalone()

		docdb_removeSourceIdentifierFromSubscriptionCmd.Flags().String("source-identifier", "", "The source identifier to be removed from the subscription, such as the instance identifier for an instance, or the name of a security group.")
		docdb_removeSourceIdentifierFromSubscriptionCmd.Flags().String("subscription-name", "", "The name of the Amazon DocumentDB event notification subscription that you want to remove a source identifier from.")
		docdb_removeSourceIdentifierFromSubscriptionCmd.MarkFlagRequired("source-identifier")
		docdb_removeSourceIdentifierFromSubscriptionCmd.MarkFlagRequired("subscription-name")
	})
	docdbCmd.AddCommand(docdb_removeSourceIdentifierFromSubscriptionCmd)
}
