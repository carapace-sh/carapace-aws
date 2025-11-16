package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_deleteDataLakeExceptionSubscriptionCmd = &cobra.Command{
	Use:   "delete-data-lake-exception-subscription",
	Short: "Deletes the specified notification subscription in Amazon Security Lake for the organization you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_deleteDataLakeExceptionSubscriptionCmd).Standalone()

	securitylakeCmd.AddCommand(securitylake_deleteDataLakeExceptionSubscriptionCmd)
}
