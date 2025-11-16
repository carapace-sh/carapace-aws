package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createAnycastIpList2020_05_31Cmd = &cobra.Command{
	Use:   "create-anycast-ip-list2020_05_31",
	Short: "Creates an Anycast static IP list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createAnycastIpList2020_05_31Cmd).Standalone()

	cloudfront_createAnycastIpList2020_05_31Cmd.Flags().String("ip-address-type", "", "The IP address type for the Anycast static IP list.")
	cloudfront_createAnycastIpList2020_05_31Cmd.Flags().String("ip-count", "", "The number of static IP addresses that are allocated to the Anycast static IP list.")
	cloudfront_createAnycastIpList2020_05_31Cmd.Flags().String("name", "", "Name of the Anycast static IP list.")
	cloudfront_createAnycastIpList2020_05_31Cmd.Flags().String("tags", "", "")
	cloudfront_createAnycastIpList2020_05_31Cmd.MarkFlagRequired("ip-count")
	cloudfront_createAnycastIpList2020_05_31Cmd.MarkFlagRequired("name")
	cloudfrontCmd.AddCommand(cloudfront_createAnycastIpList2020_05_31Cmd)
}
