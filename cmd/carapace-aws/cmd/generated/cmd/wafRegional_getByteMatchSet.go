package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_getByteMatchSetCmd = &cobra.Command{
	Use:   "get-byte-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_getByteMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_getByteMatchSetCmd).Standalone()

		wafRegional_getByteMatchSetCmd.Flags().String("byte-match-set-id", "", "The `ByteMatchSetId` of the [ByteMatchSet]() that you want to get.")
		wafRegional_getByteMatchSetCmd.MarkFlagRequired("byte-match-set-id")
	})
	wafRegionalCmd.AddCommand(wafRegional_getByteMatchSetCmd)
}
