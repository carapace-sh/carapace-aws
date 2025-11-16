package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_updateDataLakeExceptionSubscriptionCmd = &cobra.Command{
	Use:   "update-data-lake-exception-subscription",
	Short: "Updates the specified notification subscription in Amazon Security Lake for the organization you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_updateDataLakeExceptionSubscriptionCmd).Standalone()

	securitylake_updateDataLakeExceptionSubscriptionCmd.Flags().String("exception-time-to-live", "", "The time-to-live (TTL) for the exception message to remain.")
	securitylake_updateDataLakeExceptionSubscriptionCmd.Flags().String("notification-endpoint", "", "The account that is subscribed to receive exception notifications.")
	securitylake_updateDataLakeExceptionSubscriptionCmd.Flags().String("subscription-protocol", "", "The subscription protocol to which exception messages are posted.")
	securitylake_updateDataLakeExceptionSubscriptionCmd.MarkFlagRequired("notification-endpoint")
	securitylake_updateDataLakeExceptionSubscriptionCmd.MarkFlagRequired("subscription-protocol")
	securitylakeCmd.AddCommand(securitylake_updateDataLakeExceptionSubscriptionCmd)
}
