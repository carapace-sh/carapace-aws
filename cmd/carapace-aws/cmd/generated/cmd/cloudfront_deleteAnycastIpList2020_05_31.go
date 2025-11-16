package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteAnycastIpList2020_05_31Cmd = &cobra.Command{
	Use:   "delete-anycast-ip-list2020_05_31",
	Short: "Deletes an Anycast static IP list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteAnycastIpList2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_deleteAnycastIpList2020_05_31Cmd).Standalone()

		cloudfront_deleteAnycastIpList2020_05_31Cmd.Flags().String("id", "", "The ID of the Anycast static IP list.")
		cloudfront_deleteAnycastIpList2020_05_31Cmd.Flags().String("if-match", "", "The current version (`ETag` value) of the Anycast static IP list that you are deleting.")
		cloudfront_deleteAnycastIpList2020_05_31Cmd.MarkFlagRequired("id")
		cloudfront_deleteAnycastIpList2020_05_31Cmd.MarkFlagRequired("if-match")
	})
	cloudfrontCmd.AddCommand(cloudfront_deleteAnycastIpList2020_05_31Cmd)
}
