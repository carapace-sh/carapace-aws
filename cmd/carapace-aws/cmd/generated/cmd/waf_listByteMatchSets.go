package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_listByteMatchSetsCmd = &cobra.Command{
	Use:   "list-byte-match-sets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_listByteMatchSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_listByteMatchSetsCmd).Standalone()

		waf_listByteMatchSetsCmd.Flags().String("limit", "", "Specifies the number of `ByteMatchSet` objects that you want AWS WAF to return for this request.")
		waf_listByteMatchSetsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `ByteMatchSets` than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `ByteMatchSets`.")
	})
	wafCmd.AddCommand(waf_listByteMatchSetsCmd)
}
