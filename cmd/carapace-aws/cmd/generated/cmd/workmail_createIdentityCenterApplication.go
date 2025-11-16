package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_createIdentityCenterApplicationCmd = &cobra.Command{
	Use:   "create-identity-center-application",
	Short: "Creates the WorkMail application in IAM Identity Center that can be used later in the WorkMail - IdC integration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_createIdentityCenterApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_createIdentityCenterApplicationCmd).Standalone()

		workmail_createIdentityCenterApplicationCmd.Flags().String("client-token", "", "The idempotency token associated with the request.")
		workmail_createIdentityCenterApplicationCmd.Flags().String("instance-arn", "", "The Amazon Resource Name (ARN) of the instance.")
		workmail_createIdentityCenterApplicationCmd.Flags().String("name", "", "The name of the IAM Identity Center application.")
		workmail_createIdentityCenterApplicationCmd.MarkFlagRequired("instance-arn")
		workmail_createIdentityCenterApplicationCmd.MarkFlagRequired("name")
	})
	workmailCmd.AddCommand(workmail_createIdentityCenterApplicationCmd)
}
