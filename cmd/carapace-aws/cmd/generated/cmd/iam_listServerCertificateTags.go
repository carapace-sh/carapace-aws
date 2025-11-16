package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_listServerCertificateTagsCmd = &cobra.Command{
	Use:   "list-server-certificate-tags",
	Short: "Lists the tags that are attached to the specified IAM server certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_listServerCertificateTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_listServerCertificateTagsCmd).Standalone()

		iam_listServerCertificateTagsCmd.Flags().String("marker", "", "Use this parameter only when paginating results and only after you receive a response indicating that the results are truncated.")
		iam_listServerCertificateTagsCmd.Flags().String("max-items", "", "Use this only when paginating results to indicate the maximum number of items you want in the response.")
		iam_listServerCertificateTagsCmd.Flags().String("server-certificate-name", "", "The name of the IAM server certificate whose tags you want to see.")
		iam_listServerCertificateTagsCmd.MarkFlagRequired("server-certificate-name")
	})
	iamCmd.AddCommand(iam_listServerCertificateTagsCmd)
}
