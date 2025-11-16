package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_updateInstanceMetadataOptionsCmd = &cobra.Command{
	Use:   "update-instance-metadata-options",
	Short: "Modifies the Amazon Lightsail instance metadata parameters on a running or stopped instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_updateInstanceMetadataOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_updateInstanceMetadataOptionsCmd).Standalone()

		lightsail_updateInstanceMetadataOptionsCmd.Flags().String("http-endpoint", "", "Enables or disables the HTTP metadata endpoint on your instances.")
		lightsail_updateInstanceMetadataOptionsCmd.Flags().String("http-protocol-ipv6", "", "Enables or disables the IPv6 endpoint for the instance metadata service.")
		lightsail_updateInstanceMetadataOptionsCmd.Flags().String("http-put-response-hop-limit", "", "The desired HTTP PUT response hop limit for instance metadata requests.")
		lightsail_updateInstanceMetadataOptionsCmd.Flags().String("http-tokens", "", "The state of token usage for your instance metadata requests.")
		lightsail_updateInstanceMetadataOptionsCmd.Flags().String("instance-name", "", "The name of the instance for which to update metadata parameters.")
		lightsail_updateInstanceMetadataOptionsCmd.MarkFlagRequired("instance-name")
	})
	lightsailCmd.AddCommand(lightsail_updateInstanceMetadataOptionsCmd)
}
