package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_listIpsetsCmd = &cobra.Command{
	Use:   "list-ipsets",
	Short: "Retrieves an array of [IPSetSummary]() objects for the IP sets that you manage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_listIpsetsCmd).Standalone()

	wafv2_listIpsetsCmd.Flags().String("limit", "", "The maximum number of objects that you want WAF to return for this request.")
	wafv2_listIpsetsCmd.Flags().String("next-marker", "", "When you request a list of objects with a `Limit` setting, if the number of objects that are still available for retrieval exceeds the limit, WAF returns a `NextMarker` value in the response.")
	wafv2_listIpsetsCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_listIpsetsCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_listIpsetsCmd)
}
