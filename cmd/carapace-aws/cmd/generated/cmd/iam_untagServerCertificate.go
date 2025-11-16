package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_untagServerCertificateCmd = &cobra.Command{
	Use:   "untag-server-certificate",
	Short: "Removes the specified tags from the IAM server certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_untagServerCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_untagServerCertificateCmd).Standalone()

		iam_untagServerCertificateCmd.Flags().String("server-certificate-name", "", "The name of the IAM server certificate from which you want to remove tags.")
		iam_untagServerCertificateCmd.Flags().String("tag-keys", "", "A list of key names as a simple array of strings.")
		iam_untagServerCertificateCmd.MarkFlagRequired("server-certificate-name")
		iam_untagServerCertificateCmd.MarkFlagRequired("tag-keys")
	})
	iamCmd.AddCommand(iam_untagServerCertificateCmd)
}
