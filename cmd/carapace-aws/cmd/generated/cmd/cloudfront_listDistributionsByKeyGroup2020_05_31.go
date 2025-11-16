package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDistributionsByKeyGroup2020_05_31Cmd = &cobra.Command{
	Use:   "list-distributions-by-key-group2020_05_31",
	Short: "Gets a list of distribution IDs for distributions that have a cache behavior that references the specified key group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDistributionsByKeyGroup2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listDistributionsByKeyGroup2020_05_31Cmd).Standalone()

		cloudfront_listDistributionsByKeyGroup2020_05_31Cmd.Flags().String("key-group-id", "", "The ID of the key group whose associated distribution IDs you are listing.")
		cloudfront_listDistributionsByKeyGroup2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of distribution IDs.")
		cloudfront_listDistributionsByKeyGroup2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of distribution IDs that you want in the response.")
		cloudfront_listDistributionsByKeyGroup2020_05_31Cmd.MarkFlagRequired("key-group-id")
	})
	cloudfrontCmd.AddCommand(cloudfront_listDistributionsByKeyGroup2020_05_31Cmd)
}
