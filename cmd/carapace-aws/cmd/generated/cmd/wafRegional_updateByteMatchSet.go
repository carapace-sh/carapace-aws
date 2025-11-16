package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_updateByteMatchSetCmd = &cobra.Command{
	Use:   "update-byte-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_updateByteMatchSetCmd).Standalone()

	wafRegional_updateByteMatchSetCmd.Flags().String("byte-match-set-id", "", "The `ByteMatchSetId` of the [ByteMatchSet]() that you want to update.")
	wafRegional_updateByteMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_updateByteMatchSetCmd.Flags().String("updates", "", "An array of `ByteMatchSetUpdate` objects that you want to insert into or delete from a [ByteMatchSet]().")
	wafRegional_updateByteMatchSetCmd.MarkFlagRequired("byte-match-set-id")
	wafRegional_updateByteMatchSetCmd.MarkFlagRequired("change-token")
	wafRegional_updateByteMatchSetCmd.MarkFlagRequired("updates")
	wafRegionalCmd.AddCommand(wafRegional_updateByteMatchSetCmd)
}
