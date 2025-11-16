package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_tagServerCertificateCmd = &cobra.Command{
	Use:   "tag-server-certificate",
	Short: "Adds one or more tags to an IAM server certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_tagServerCertificateCmd).Standalone()

	iam_tagServerCertificateCmd.Flags().String("server-certificate-name", "", "The name of the IAM server certificate to which you want to add tags.")
	iam_tagServerCertificateCmd.Flags().String("tags", "", "The list of tags that you want to attach to the IAM server certificate.")
	iam_tagServerCertificateCmd.MarkFlagRequired("server-certificate-name")
	iam_tagServerCertificateCmd.MarkFlagRequired("tags")
	iamCmd.AddCommand(iam_tagServerCertificateCmd)
}
