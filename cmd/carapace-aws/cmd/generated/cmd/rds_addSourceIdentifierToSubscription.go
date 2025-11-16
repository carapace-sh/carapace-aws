package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_addSourceIdentifierToSubscriptionCmd = &cobra.Command{
	Use:   "add-source-identifier-to-subscription",
	Short: "Adds a source identifier to an existing RDS event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_addSourceIdentifierToSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_addSourceIdentifierToSubscriptionCmd).Standalone()

		rds_addSourceIdentifierToSubscriptionCmd.Flags().String("source-identifier", "", "The identifier of the event source to be added.")
		rds_addSourceIdentifierToSubscriptionCmd.Flags().String("subscription-name", "", "The name of the RDS event notification subscription you want to add a source identifier to.")
		rds_addSourceIdentifierToSubscriptionCmd.MarkFlagRequired("source-identifier")
		rds_addSourceIdentifierToSubscriptionCmd.MarkFlagRequired("subscription-name")
	})
	rdsCmd.AddCommand(rds_addSourceIdentifierToSubscriptionCmd)
}
