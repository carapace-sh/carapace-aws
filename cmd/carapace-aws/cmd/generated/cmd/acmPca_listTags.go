package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acmPca_listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "Lists the tags, if any, that are associated with your private CA or one that has been shared with you.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acmPca_listTagsCmd).Standalone()

	acmPca_listTagsCmd.Flags().String("certificate-authority-arn", "", "The Amazon Resource Name (ARN) that was returned when you called the [CreateCertificateAuthority](https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html) action.")
	acmPca_listTagsCmd.Flags().String("max-results", "", "Use this parameter when paginating results to specify the maximum number of items to return in the response.")
	acmPca_listTagsCmd.Flags().String("next-token", "", "Use this parameter when paginating results in a subsequent request after you receive a response with truncated results.")
	acmPca_listTagsCmd.MarkFlagRequired("certificate-authority-arn")
	acmPcaCmd.AddCommand(acmPca_listTagsCmd)
}
