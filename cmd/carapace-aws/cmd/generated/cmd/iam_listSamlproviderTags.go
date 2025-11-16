package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listSamlproviderTagsCmd = &cobra.Command{
	Use:   "list-samlprovider-tags",
	Short: "Lists the tags that are attached to the specified Security Assertion Markup Language (SAML) identity provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listSamlproviderTagsCmd).Standalone()

	iam_listSamlproviderTagsCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
	iam_listSamlproviderTagsCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
	iam_listSamlproviderTagsCmd.Flags().String("samlprovider-arn", "", "The ARN of the Security Assertion Markup Language (SAML) identity provider whose tags you want to see.")
	iam_listSamlproviderTagsCmd.MarkFlagRequired("samlprovider-arn")
	iamCmd.AddCommand(iam_listSamlproviderTagsCmd)
}
