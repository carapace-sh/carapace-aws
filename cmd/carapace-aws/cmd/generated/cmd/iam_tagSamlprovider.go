package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_tagSamlproviderCmd = &cobra.Command{
	Use:   "tag-samlprovider",
	Short: "Adds one or more tags to a Security Assertion Markup Language (SAML) identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_tagSamlproviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_tagSamlproviderCmd).Standalone()

		iam_tagSamlproviderCmd.Flags().String("samlprovider-arn", "", "The ARN of the SAML identity provider in IAM to which you want to add tags.")
		iam_tagSamlproviderCmd.Flags().String("tags", "", "The list of tags that you want to attach to the SAML identity provider in IAM.")
		iam_tagSamlproviderCmd.MarkFlagRequired("samlprovider-arn")
		iam_tagSamlproviderCmd.MarkFlagRequired("tags")
	})
	iamCmd.AddCommand(iam_tagSamlproviderCmd)
}
