package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listByteMatchSetsCmd = &cobra.Command{
	Use:   "list-byte-match-sets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listByteMatchSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_listByteMatchSetsCmd).Standalone()

		wafRegional_listByteMatchSetsCmd.Flags().String("limit", "", "Specifies the number of `ByteMatchSet` objects that you want AWS WAF to return for this request.")
		wafRegional_listByteMatchSetsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `ByteMatchSets` than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `ByteMatchSets`.")
	})
	wafRegionalCmd.AddCommand(wafRegional_listByteMatchSetsCmd)
}
