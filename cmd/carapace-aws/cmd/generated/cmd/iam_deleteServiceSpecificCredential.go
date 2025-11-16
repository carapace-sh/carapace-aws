package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteServiceSpecificCredentialCmd = &cobra.Command{
	Use:   "delete-service-specific-credential",
	Short: "Deletes the specified service-specific credential.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteServiceSpecificCredentialCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_deleteServiceSpecificCredentialCmd).Standalone()

		iam_deleteServiceSpecificCredentialCmd.Flags().String("service-specific-credential-id", "", "The unique identifier of the service-specific credential.")
		iam_deleteServiceSpecificCredentialCmd.Flags().String("user-name", "", "The name of the IAM user associated with the service-specific credential.")
		iam_deleteServiceSpecificCredentialCmd.MarkFlagRequired("service-specific-credential-id")
	})
	iamCmd.AddCommand(iam_deleteServiceSpecificCredentialCmd)
}
