package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listCachePolicies2020_05_31Cmd = &cobra.Command{
	Use:   "list-cache-policies2020_05_31",
	Short: "Gets a list of cache policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listCachePolicies2020_05_31Cmd).Standalone()

	cloudfront_listCachePolicies2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of cache policies.")
	cloudfront_listCachePolicies2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of cache policies that you want in the response.")
	cloudfront_listCachePolicies2020_05_31Cmd.Flags().String("type", "", "A filter to return only the specified kinds of cache policies.")
	cloudfrontCmd.AddCommand(cloudfront_listCachePolicies2020_05_31Cmd)
}
