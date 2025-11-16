package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getByteMatchSetCmd = &cobra.Command{
	Use:   "get-byte-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getByteMatchSetCmd).Standalone()

	waf_getByteMatchSetCmd.Flags().String("byte-match-set-id", "", "The `ByteMatchSetId` of the [ByteMatchSet]() that you want to get.")
	waf_getByteMatchSetCmd.MarkFlagRequired("byte-match-set-id")
	wafCmd.AddCommand(waf_getByteMatchSetCmd)
}
