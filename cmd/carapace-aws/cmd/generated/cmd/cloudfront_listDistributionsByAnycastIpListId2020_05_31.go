package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDistributionsByAnycastIpListId2020_05_31Cmd = &cobra.Command{
	Use:   "list-distributions-by-anycast-ip-list-id2020_05_31",
	Short: "Lists the distributions in your account that are associated with the specified `AnycastIpListId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDistributionsByAnycastIpListId2020_05_31Cmd).Standalone()

	cloudfront_listDistributionsByAnycastIpListId2020_05_31Cmd.Flags().String("anycast-ip-list-id", "", "The ID of the Anycast static IP list.")
	cloudfront_listDistributionsByAnycastIpListId2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list.")
	cloudfront_listDistributionsByAnycastIpListId2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of distributions that you want returned in the response.")
	cloudfront_listDistributionsByAnycastIpListId2020_05_31Cmd.MarkFlagRequired("anycast-ip-list-id")
	cloudfrontCmd.AddCommand(cloudfront_listDistributionsByAnycastIpListId2020_05_31Cmd)
}
