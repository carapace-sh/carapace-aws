package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getAnycastIpList2020_05_31Cmd = &cobra.Command{
	Use:   "get-anycast-ip-list2020_05_31",
	Short: "Gets an Anycast static IP list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getAnycastIpList2020_05_31Cmd).Standalone()

	cloudfront_getAnycastIpList2020_05_31Cmd.Flags().String("id", "", "The ID of the Anycast static IP list.")
	cloudfront_getAnycastIpList2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_getAnycastIpList2020_05_31Cmd)
}
