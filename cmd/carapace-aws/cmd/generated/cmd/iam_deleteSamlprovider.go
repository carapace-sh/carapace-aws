package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteSamlproviderCmd = &cobra.Command{
	Use:   "delete-samlprovider",
	Short: "Deletes a SAML provider resource in IAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteSamlproviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_deleteSamlproviderCmd).Standalone()

		iam_deleteSamlproviderCmd.Flags().String("samlprovider-arn", "", "The Amazon Resource Name (ARN) of the SAML provider to delete.")
		iam_deleteSamlproviderCmd.MarkFlagRequired("samlprovider-arn")
	})
	iamCmd.AddCommand(iam_deleteSamlproviderCmd)
}
