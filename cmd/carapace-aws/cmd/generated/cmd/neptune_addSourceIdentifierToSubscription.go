package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_addSourceIdentifierToSubscriptionCmd = &cobra.Command{
	Use:   "add-source-identifier-to-subscription",
	Short: "Adds a source identifier to an existing event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_addSourceIdentifierToSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_addSourceIdentifierToSubscriptionCmd).Standalone()

		neptune_addSourceIdentifierToSubscriptionCmd.Flags().String("source-identifier", "", "The identifier of the event source to be added.")
		neptune_addSourceIdentifierToSubscriptionCmd.Flags().String("subscription-name", "", "The name of the event notification subscription you want to add a source identifier to.")
		neptune_addSourceIdentifierToSubscriptionCmd.MarkFlagRequired("source-identifier")
		neptune_addSourceIdentifierToSubscriptionCmd.MarkFlagRequired("subscription-name")
	})
	neptuneCmd.AddCommand(neptune_addSourceIdentifierToSubscriptionCmd)
}
