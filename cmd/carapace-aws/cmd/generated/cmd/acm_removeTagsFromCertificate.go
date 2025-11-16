package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_removeTagsFromCertificateCmd = &cobra.Command{
	Use:   "remove-tags-from-certificate",
	Short: "Remove one or more tags from an ACM certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_removeTagsFromCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acm_removeTagsFromCertificateCmd).Standalone()

		acm_removeTagsFromCertificateCmd.Flags().String("certificate-arn", "", "String that contains the ARN of the ACM Certificate with one or more tags that you want to remove.")
		acm_removeTagsFromCertificateCmd.Flags().String("tags", "", "The key-value pair that defines the tag to remove.")
		acm_removeTagsFromCertificateCmd.MarkFlagRequired("certificate-arn")
		acm_removeTagsFromCertificateCmd.MarkFlagRequired("tags")
	})
	acmCmd.AddCommand(acm_removeTagsFromCertificateCmd)
}
