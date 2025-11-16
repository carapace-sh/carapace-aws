package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_updateDistributionCmd = &cobra.Command{
	Use:   "update-distribution",
	Short: "Updates an existing Amazon Lightsail content delivery network (CDN) distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_updateDistributionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_updateDistributionCmd).Standalone()

		lightsail_updateDistributionCmd.Flags().String("cache-behavior-settings", "", "An object that describes the cache behavior settings for the distribution.")
		lightsail_updateDistributionCmd.Flags().String("cache-behaviors", "", "An array of objects that describe the per-path cache behavior for the distribution.")
		lightsail_updateDistributionCmd.Flags().String("certificate-name", "", "The name of the SSL/TLS certificate that you want to attach to the distribution.")
		lightsail_updateDistributionCmd.Flags().String("default-cache-behavior", "", "An object that describes the default cache behavior for the distribution.")
		lightsail_updateDistributionCmd.Flags().String("distribution-name", "", "The name of the distribution to update.")
		lightsail_updateDistributionCmd.Flags().String("is-enabled", "", "Indicates whether to enable the distribution.")
		lightsail_updateDistributionCmd.Flags().String("origin", "", "An object that describes the origin resource for the distribution, such as a Lightsail instance, bucket, or load balancer.")
		lightsail_updateDistributionCmd.Flags().String("use-default-certificate", "", "Indicates whether the default SSL/TLS certificate is attached to the distribution.")
		lightsail_updateDistributionCmd.Flags().String("viewer-minimum-tls-protocol-version", "", "Use this parameter to update the minimum TLS protocol version for the SSL/TLS certificate that's attached to the distribution.")
		lightsail_updateDistributionCmd.MarkFlagRequired("distribution-name")
	})
	lightsailCmd.AddCommand(lightsail_updateDistributionCmd)
}
