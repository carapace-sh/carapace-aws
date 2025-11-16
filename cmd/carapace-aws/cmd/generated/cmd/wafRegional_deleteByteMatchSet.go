package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteByteMatchSetCmd = &cobra.Command{
	Use:   "delete-byte-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteByteMatchSetCmd).Standalone()

	wafRegional_deleteByteMatchSetCmd.Flags().String("byte-match-set-id", "", "The `ByteMatchSetId` of the [ByteMatchSet]() that you want to delete.")
	wafRegional_deleteByteMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_deleteByteMatchSetCmd.MarkFlagRequired("byte-match-set-id")
	wafRegional_deleteByteMatchSetCmd.MarkFlagRequired("change-token")
	wafRegionalCmd.AddCommand(wafRegional_deleteByteMatchSetCmd)
}
