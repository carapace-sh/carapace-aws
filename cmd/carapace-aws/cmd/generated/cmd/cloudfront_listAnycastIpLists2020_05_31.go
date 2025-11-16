package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listAnycastIpLists2020_05_31Cmd = &cobra.Command{
	Use:   "list-anycast-ip-lists2020_05_31",
	Short: "Lists your Anycast static IP lists.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listAnycastIpLists2020_05_31Cmd).Standalone()

	cloudfront_listAnycastIpLists2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list.")
	cloudfront_listAnycastIpLists2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of Anycast static IP lists that you want returned in the response.")
	cloudfrontCmd.AddCommand(cloudfront_listAnycastIpLists2020_05_31Cmd)
}
