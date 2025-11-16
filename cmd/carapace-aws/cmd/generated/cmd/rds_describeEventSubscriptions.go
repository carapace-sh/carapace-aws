package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeEventSubscriptionsCmd = &cobra.Command{
	Use:   "describe-event-subscriptions",
	Short: "Lists all the subscription descriptions for a customer account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeEventSubscriptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeEventSubscriptionsCmd).Standalone()

		rds_describeEventSubscriptionsCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
		rds_describeEventSubscriptionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous DescribeOrderableDBInstanceOptions request.")
		rds_describeEventSubscriptionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeEventSubscriptionsCmd.Flags().String("subscription-name", "", "The name of the RDS event notification subscription you want to describe.")
	})
	rdsCmd.AddCommand(rds_describeEventSubscriptionsCmd)
}
