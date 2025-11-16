package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteIdentityCenterApplicationCmd = &cobra.Command{
	Use:   "delete-identity-center-application",
	Short: "Deletes the IAM Identity Center application from WorkMail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteIdentityCenterApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_deleteIdentityCenterApplicationCmd).Standalone()

		workmail_deleteIdentityCenterApplicationCmd.Flags().String("application-arn", "", "The Amazon Resource Name (ARN) of the application.")
		workmail_deleteIdentityCenterApplicationCmd.MarkFlagRequired("application-arn")
	})
	workmailCmd.AddCommand(workmail_deleteIdentityCenterApplicationCmd)
}
