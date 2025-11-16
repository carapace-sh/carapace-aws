package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateAnycastIpList2020_05_31Cmd = &cobra.Command{
	Use:   "update-anycast-ip-list2020_05_31",
	Short: "Updates an Anycast static IP list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateAnycastIpList2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_updateAnycastIpList2020_05_31Cmd).Standalone()

		cloudfront_updateAnycastIpList2020_05_31Cmd.Flags().String("id", "", "The ID of the Anycast static IP list.")
		cloudfront_updateAnycastIpList2020_05_31Cmd.Flags().String("if-match", "", "The current version (ETag value) of the Anycast static IP list that you are updating.")
		cloudfront_updateAnycastIpList2020_05_31Cmd.Flags().String("ip-address-type", "", "The IP address type for the Anycast static IP list.")
		cloudfront_updateAnycastIpList2020_05_31Cmd.MarkFlagRequired("id")
		cloudfront_updateAnycastIpList2020_05_31Cmd.MarkFlagRequired("if-match")
	})
	cloudfrontCmd.AddCommand(cloudfront_updateAnycastIpList2020_05_31Cmd)
}
