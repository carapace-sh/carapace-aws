package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateConnectionGroup2020_05_31Cmd = &cobra.Command{
	Use:   "update-connection-group2020_05_31",
	Short: "Updates a connection group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateConnectionGroup2020_05_31Cmd).Standalone()

	cloudfront_updateConnectionGroup2020_05_31Cmd.Flags().String("anycast-ip-list-id", "", "The ID of the Anycast static IP list.")
	cloudfront_updateConnectionGroup2020_05_31Cmd.Flags().String("enabled", "", "Whether the connection group is enabled.")
	cloudfront_updateConnectionGroup2020_05_31Cmd.Flags().String("id", "", "The ID of the connection group.")
	cloudfront_updateConnectionGroup2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the connection group that you're updating.")
	cloudfront_updateConnectionGroup2020_05_31Cmd.Flags().String("ipv6-enabled", "", "Enable IPv6 for the connection group.")
	cloudfront_updateConnectionGroup2020_05_31Cmd.MarkFlagRequired("id")
	cloudfront_updateConnectionGroup2020_05_31Cmd.MarkFlagRequired("if-match")
	cloudfrontCmd.AddCommand(cloudfront_updateConnectionGroup2020_05_31Cmd)
}
