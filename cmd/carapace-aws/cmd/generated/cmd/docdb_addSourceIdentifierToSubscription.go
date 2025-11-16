package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_addSourceIdentifierToSubscriptionCmd = &cobra.Command{
	Use:   "add-source-identifier-to-subscription",
	Short: "Adds a source identifier to an existing event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_addSourceIdentifierToSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_addSourceIdentifierToSubscriptionCmd).Standalone()

		docdb_addSourceIdentifierToSubscriptionCmd.Flags().String("source-identifier", "", "The identifier of the event source to be added:")
		docdb_addSourceIdentifierToSubscriptionCmd.Flags().String("subscription-name", "", "The name of the Amazon DocumentDB event notification subscription that you want to add a source identifier to.")
		docdb_addSourceIdentifierToSubscriptionCmd.MarkFlagRequired("source-identifier")
		docdb_addSourceIdentifierToSubscriptionCmd.MarkFlagRequired("subscription-name")
	})
	docdbCmd.AddCommand(docdb_addSourceIdentifierToSubscriptionCmd)
}
