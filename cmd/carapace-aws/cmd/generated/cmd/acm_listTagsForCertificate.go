package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_listTagsForCertificateCmd = &cobra.Command{
	Use:   "list-tags-for-certificate",
	Short: "Lists the tags that have been applied to the ACM certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_listTagsForCertificateCmd).Standalone()

	acm_listTagsForCertificateCmd.Flags().String("certificate-arn", "", "String that contains the ARN of the ACM certificate for which you want to list the tags.")
	acm_listTagsForCertificateCmd.MarkFlagRequired("certificate-arn")
	acmCmd.AddCommand(acm_listTagsForCertificateCmd)
}
