package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createConnectionGroup2020_05_31Cmd = &cobra.Command{
	Use:   "create-connection-group2020_05_31",
	Short: "Creates a connection group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createConnectionGroup2020_05_31Cmd).Standalone()

	cloudfront_createConnectionGroup2020_05_31Cmd.Flags().String("anycast-ip-list-id", "", "The ID of the Anycast static IP list.")
	cloudfront_createConnectionGroup2020_05_31Cmd.Flags().String("enabled", "", "Enable the connection group.")
	cloudfront_createConnectionGroup2020_05_31Cmd.Flags().String("ipv6-enabled", "", "Enable IPv6 for the connection group.")
	cloudfront_createConnectionGroup2020_05_31Cmd.Flags().String("name", "", "The name of the connection group.")
	cloudfront_createConnectionGroup2020_05_31Cmd.Flags().String("tags", "", "")
	cloudfront_createConnectionGroup2020_05_31Cmd.MarkFlagRequired("name")
	cloudfrontCmd.AddCommand(cloudfront_createConnectionGroup2020_05_31Cmd)
}
