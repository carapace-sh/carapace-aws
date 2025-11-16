package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeEventSubscriptionsCmd = &cobra.Command{
	Use:   "describe-event-subscriptions",
	Short: "Lists descriptions of all the Amazon Redshift event notification subscriptions for a customer account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeEventSubscriptionsCmd).Standalone()

	redshift_describeEventSubscriptionsCmd.Flags().String("marker", "", "An optional parameter that specifies the starting point to return a set of response records.")
	redshift_describeEventSubscriptionsCmd.Flags().String("max-records", "", "The maximum number of response records to return in each call.")
	redshift_describeEventSubscriptionsCmd.Flags().String("subscription-name", "", "The name of the Amazon Redshift event notification subscription to be described.")
	redshift_describeEventSubscriptionsCmd.Flags().String("tag-keys", "", "A tag key or keys for which you want to return all matching event notification subscriptions that are associated with the specified key or keys.")
	redshift_describeEventSubscriptionsCmd.Flags().String("tag-values", "", "A tag value or values for which you want to return all matching event notification subscriptions that are associated with the specified tag value or values.")
	redshiftCmd.AddCommand(redshift_describeEventSubscriptionsCmd)
}
