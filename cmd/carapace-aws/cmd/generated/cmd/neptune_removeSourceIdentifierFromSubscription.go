package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_removeSourceIdentifierFromSubscriptionCmd = &cobra.Command{
	Use:   "remove-source-identifier-from-subscription",
	Short: "Removes a source identifier from an existing event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_removeSourceIdentifierFromSubscriptionCmd).Standalone()

	neptune_removeSourceIdentifierFromSubscriptionCmd.Flags().String("source-identifier", "", "The source identifier to be removed from the subscription, such as the **DB instance identifier** for a DB instance or the name of a security group.")
	neptune_removeSourceIdentifierFromSubscriptionCmd.Flags().String("subscription-name", "", "The name of the event notification subscription you want to remove a source identifier from.")
	neptune_removeSourceIdentifierFromSubscriptionCmd.MarkFlagRequired("source-identifier")
	neptune_removeSourceIdentifierFromSubscriptionCmd.MarkFlagRequired("subscription-name")
	neptuneCmd.AddCommand(neptune_removeSourceIdentifierFromSubscriptionCmd)
}
