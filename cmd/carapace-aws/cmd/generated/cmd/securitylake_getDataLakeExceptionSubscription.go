package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securitylake_getDataLakeExceptionSubscriptionCmd = &cobra.Command{
	Use:   "get-data-lake-exception-subscription",
	Short: "Retrieves the protocol and endpoint that were provided when subscribing to Amazon SNS topics for exception notifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securitylake_getDataLakeExceptionSubscriptionCmd).Standalone()

	securitylakeCmd.AddCommand(securitylake_getDataLakeExceptionSubscriptionCmd)
}
