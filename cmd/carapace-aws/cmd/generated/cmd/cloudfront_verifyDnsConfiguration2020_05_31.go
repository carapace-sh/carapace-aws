package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_verifyDnsConfiguration2020_05_31Cmd = &cobra.Command{
	Use:   "verify-dns-configuration2020_05_31",
	Short: "Verify the DNS configuration for your domain names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_verifyDnsConfiguration2020_05_31Cmd).Standalone()

	cloudfront_verifyDnsConfiguration2020_05_31Cmd.Flags().String("domain", "", "The domain name that you're verifying.")
	cloudfront_verifyDnsConfiguration2020_05_31Cmd.Flags().String("identifier", "", "The identifier of the distribution tenant.")
	cloudfront_verifyDnsConfiguration2020_05_31Cmd.MarkFlagRequired("identifier")
	cloudfrontCmd.AddCommand(cloudfront_verifyDnsConfiguration2020_05_31Cmd)
}
