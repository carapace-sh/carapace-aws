package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createDistributionCmd = &cobra.Command{
	Use:   "create-distribution",
	Short: "Creates an Amazon Lightsail content delivery network (CDN) distribution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createDistributionCmd).Standalone()

	lightsail_createDistributionCmd.Flags().String("bundle-id", "", "The bundle ID to use for the distribution.")
	lightsail_createDistributionCmd.Flags().String("cache-behavior-settings", "", "An object that describes the cache behavior settings for the distribution.")
	lightsail_createDistributionCmd.Flags().String("cache-behaviors", "", "An array of objects that describe the per-path cache behavior for the distribution.")
	lightsail_createDistributionCmd.Flags().String("certificate-name", "", "The name of the SSL/TLS certificate that you want to attach to the distribution.")
	lightsail_createDistributionCmd.Flags().String("default-cache-behavior", "", "An object that describes the default cache behavior for the distribution.")
	lightsail_createDistributionCmd.Flags().String("distribution-name", "", "The name for the distribution.")
	lightsail_createDistributionCmd.Flags().String("ip-address-type", "", "The IP address type for the distribution.")
	lightsail_createDistributionCmd.Flags().String("origin", "", "An object that describes the origin resource for the distribution, such as a Lightsail instance, bucket, or load balancer.")
	lightsail_createDistributionCmd.Flags().String("tags", "", "The tag keys and optional values to add to the distribution during create.")
	lightsail_createDistributionCmd.Flags().String("viewer-minimum-tls-protocol-version", "", "The minimum TLS protocol version for the SSL/TLS certificate.")
	lightsail_createDistributionCmd.MarkFlagRequired("bundle-id")
	lightsail_createDistributionCmd.MarkFlagRequired("default-cache-behavior")
	lightsail_createDistributionCmd.MarkFlagRequired("distribution-name")
	lightsail_createDistributionCmd.MarkFlagRequired("origin")
	lightsailCmd.AddCommand(lightsail_createDistributionCmd)
}
