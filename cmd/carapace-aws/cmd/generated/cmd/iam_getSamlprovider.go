package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getSamlproviderCmd = &cobra.Command{
	Use:   "get-samlprovider",
	Short: "Returns the SAML provider metadocument that was uploaded when the IAM SAML provider resource object was created or updated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getSamlproviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_getSamlproviderCmd).Standalone()

		iam_getSamlproviderCmd.Flags().String("samlprovider-arn", "", "The Amazon Resource Name (ARN) of the SAML provider resource object in IAM to get information about.")
		iam_getSamlproviderCmd.MarkFlagRequired("samlprovider-arn")
	})
	iamCmd.AddCommand(iam_getSamlproviderCmd)
}
