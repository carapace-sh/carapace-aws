package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteByteMatchSetCmd = &cobra.Command{
	Use:   "delete-byte-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteByteMatchSetCmd).Standalone()

	waf_deleteByteMatchSetCmd.Flags().String("byte-match-set-id", "", "The `ByteMatchSetId` of the [ByteMatchSet]() that you want to delete.")
	waf_deleteByteMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	waf_deleteByteMatchSetCmd.MarkFlagRequired("byte-match-set-id")
	waf_deleteByteMatchSetCmd.MarkFlagRequired("change-token")
	wafCmd.AddCommand(waf_deleteByteMatchSetCmd)
}
