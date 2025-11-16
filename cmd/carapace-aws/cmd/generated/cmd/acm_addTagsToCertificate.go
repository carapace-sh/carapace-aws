package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_addTagsToCertificateCmd = &cobra.Command{
	Use:   "add-tags-to-certificate",
	Short: "Adds one or more tags to an ACM certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_addTagsToCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acm_addTagsToCertificateCmd).Standalone()

		acm_addTagsToCertificateCmd.Flags().String("certificate-arn", "", "String that contains the ARN of the ACM certificate to which the tag is to be applied.")
		acm_addTagsToCertificateCmd.Flags().String("tags", "", "The key-value pair that defines the tag.")
		acm_addTagsToCertificateCmd.MarkFlagRequired("certificate-arn")
		acm_addTagsToCertificateCmd.MarkFlagRequired("tags")
	})
	acmCmd.AddCommand(acm_addTagsToCertificateCmd)
}
