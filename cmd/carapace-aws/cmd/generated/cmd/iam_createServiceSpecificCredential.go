package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createServiceSpecificCredentialCmd = &cobra.Command{
	Use:   "create-service-specific-credential",
	Short: "Generates a set of credentials consisting of a user name and password that can be used to access the service specified in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createServiceSpecificCredentialCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_createServiceSpecificCredentialCmd).Standalone()

		iam_createServiceSpecificCredentialCmd.Flags().String("credential-age-days", "", "The number of days until the service specific credential expires.")
		iam_createServiceSpecificCredentialCmd.Flags().String("service-name", "", "The name of the Amazon Web Services service that is to be associated with the credentials.")
		iam_createServiceSpecificCredentialCmd.Flags().String("user-name", "", "The name of the IAM user that is to be associated with the credentials.")
		iam_createServiceSpecificCredentialCmd.MarkFlagRequired("service-name")
		iam_createServiceSpecificCredentialCmd.MarkFlagRequired("user-name")
	})
	iamCmd.AddCommand(iam_createServiceSpecificCredentialCmd)
}
