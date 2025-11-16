package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_untagSamlproviderCmd = &cobra.Command{
	Use:   "untag-samlprovider",
	Short: "Removes the specified tags from the specified Security Assertion Markup Language (SAML) identity provider in IAM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_untagSamlproviderCmd).Standalone()

	iam_untagSamlproviderCmd.Flags().String("samlprovider-arn", "", "The ARN of the SAML identity provider in IAM from which you want to remove tags.")
	iam_untagSamlproviderCmd.Flags().String("tag-keys", "", "A list of key names as a simple array of strings.")
	iam_untagSamlproviderCmd.MarkFlagRequired("samlprovider-arn")
	iam_untagSamlproviderCmd.MarkFlagRequired("tag-keys")
	iamCmd.AddCommand(iam_untagSamlproviderCmd)
}
