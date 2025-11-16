package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteEventSubscriptionCmd = &cobra.Command{
	Use:   "delete-event-subscription",
	Short: "Deletes an RDS event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteEventSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteEventSubscriptionCmd).Standalone()

		rds_deleteEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the RDS event notification subscription you want to delete.")
		rds_deleteEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	})
	rdsCmd.AddCommand(rds_deleteEventSubscriptionCmd)
}
