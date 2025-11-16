package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getManagedCertificateDetails2020_05_31Cmd = &cobra.Command{
	Use:   "get-managed-certificate-details2020_05_31",
	Short: "Gets details about the CloudFront managed ACM certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getManagedCertificateDetails2020_05_31Cmd).Standalone()

	cloudfront_getManagedCertificateDetails2020_05_31Cmd.Flags().String("identifier", "", "The identifier of the distribution tenant.")
	cloudfront_getManagedCertificateDetails2020_05_31Cmd.MarkFlagRequired("identifier")
	cloudfrontCmd.AddCommand(cloudfront_getManagedCertificateDetails2020_05_31Cmd)
}
