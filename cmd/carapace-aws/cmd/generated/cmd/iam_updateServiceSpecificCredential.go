package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_updateServiceSpecificCredentialCmd = &cobra.Command{
	Use:   "update-service-specific-credential",
	Short: "Sets the status of a service-specific credential to `Active` or `Inactive`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_updateServiceSpecificCredentialCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_updateServiceSpecificCredentialCmd).Standalone()

		iam_updateServiceSpecificCredentialCmd.Flags().String("service-specific-credential-id", "", "The unique identifier of the service-specific credential.")
		iam_updateServiceSpecificCredentialCmd.Flags().String("status", "", "The status to be assigned to the service-specific credential.")
		iam_updateServiceSpecificCredentialCmd.Flags().String("user-name", "", "The name of the IAM user associated with the service-specific credential.")
		iam_updateServiceSpecificCredentialCmd.MarkFlagRequired("service-specific-credential-id")
		iam_updateServiceSpecificCredentialCmd.MarkFlagRequired("status")
	})
	iamCmd.AddCommand(iam_updateServiceSpecificCredentialCmd)
}
