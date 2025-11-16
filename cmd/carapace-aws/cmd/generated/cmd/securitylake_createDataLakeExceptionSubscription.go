package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_createDataLakeExceptionSubscriptionCmd = &cobra.Command{
	Use:   "create-data-lake-exception-subscription",
	Short: "Creates the specified notification subscription in Amazon Security Lake for the organization you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_createDataLakeExceptionSubscriptionCmd).Standalone()

	securitylake_createDataLakeExceptionSubscriptionCmd.Flags().String("exception-time-to-live", "", "The expiration period and time-to-live (TTL).")
	securitylake_createDataLakeExceptionSubscriptionCmd.Flags().String("notification-endpoint", "", "The Amazon Web Services account where you want to receive exception notifications.")
	securitylake_createDataLakeExceptionSubscriptionCmd.Flags().String("subscription-protocol", "", "The subscription protocol to which exception notifications are posted.")
	securitylake_createDataLakeExceptionSubscriptionCmd.MarkFlagRequired("notification-endpoint")
	securitylake_createDataLakeExceptionSubscriptionCmd.MarkFlagRequired("subscription-protocol")
	securitylakeCmd.AddCommand(securitylake_createDataLakeExceptionSubscriptionCmd)
}
