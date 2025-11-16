package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createDelegationRequestCmd = &cobra.Command{
	Use:   "create-delegation-request",
	Short: "This API is currently unavailable for general use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createDelegationRequestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_createDelegationRequestCmd).Standalone()

		iam_createDelegationRequestCmd.Flags().String("description", "", "")
		iam_createDelegationRequestCmd.Flags().String("notification-channel", "", "")
		iam_createDelegationRequestCmd.Flags().String("only-send-by-owner", "", "")
		iam_createDelegationRequestCmd.Flags().String("owner-account-id", "", "")
		iam_createDelegationRequestCmd.Flags().String("permissions", "", "")
		iam_createDelegationRequestCmd.Flags().String("redirect-url", "", "")
		iam_createDelegationRequestCmd.Flags().String("request-message", "", "")
		iam_createDelegationRequestCmd.Flags().String("requestor-workflow-id", "", "")
		iam_createDelegationRequestCmd.Flags().String("session-duration", "", "")
		iam_createDelegationRequestCmd.MarkFlagRequired("description")
		iam_createDelegationRequestCmd.MarkFlagRequired("notification-channel")
		iam_createDelegationRequestCmd.MarkFlagRequired("permissions")
		iam_createDelegationRequestCmd.MarkFlagRequired("requestor-workflow-id")
		iam_createDelegationRequestCmd.MarkFlagRequired("session-duration")
	})
	iamCmd.AddCommand(iam_createDelegationRequestCmd)
}
