package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_removeSourceIdentifierFromSubscriptionCmd = &cobra.Command{
	Use:   "remove-source-identifier-from-subscription",
	Short: "Removes a source identifier from an existing RDS event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_removeSourceIdentifierFromSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_removeSourceIdentifierFromSubscriptionCmd).Standalone()

		rds_removeSourceIdentifierFromSubscriptionCmd.Flags().String("source-identifier", "", "The source identifier to be removed from the subscription, such as the **DB instance identifier** for a DB instance or the name of a security group.")
		rds_removeSourceIdentifierFromSubscriptionCmd.Flags().String("subscription-name", "", "The name of the RDS event notification subscription you want to remove a source identifier from.")
		rds_removeSourceIdentifierFromSubscriptionCmd.MarkFlagRequired("source-identifier")
		rds_removeSourceIdentifierFromSubscriptionCmd.MarkFlagRequired("subscription-name")
	})
	rdsCmd.AddCommand(rds_removeSourceIdentifierFromSubscriptionCmd)
}
