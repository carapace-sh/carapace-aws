package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteEventSubscriptionCmd = &cobra.Command{
	Use:   "delete-event-subscription",
	Short: "Deletes an Amazon Redshift event notification subscription.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteEventSubscriptionCmd).Standalone()

	redshift_deleteEventSubscriptionCmd.Flags().String("subscription-name", "", "The name of the Amazon Redshift event notification subscription to be deleted.")
	redshift_deleteEventSubscriptionCmd.MarkFlagRequired("subscription-name")
	redshiftCmd.AddCommand(redshift_deleteEventSubscriptionCmd)
}
