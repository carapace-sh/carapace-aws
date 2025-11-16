package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_resetServiceSpecificCredentialCmd = &cobra.Command{
	Use:   "reset-service-specific-credential",
	Short: "Resets the password for a service-specific credential.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_resetServiceSpecificCredentialCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_resetServiceSpecificCredentialCmd).Standalone()

		iam_resetServiceSpecificCredentialCmd.Flags().String("service-specific-credential-id", "", "The unique identifier of the service-specific credential.")
		iam_resetServiceSpecificCredentialCmd.Flags().String("user-name", "", "The name of the IAM user associated with the service-specific credential.")
		iam_resetServiceSpecificCredentialCmd.MarkFlagRequired("service-specific-credential-id")
	})
	iamCmd.AddCommand(iam_resetServiceSpecificCredentialCmd)
}
