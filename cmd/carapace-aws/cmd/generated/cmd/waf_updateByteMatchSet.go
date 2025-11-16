package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_updateByteMatchSetCmd = &cobra.Command{
	Use:   "update-byte-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_updateByteMatchSetCmd).Standalone()

	waf_updateByteMatchSetCmd.Flags().String("byte-match-set-id", "", "The `ByteMatchSetId` of the [ByteMatchSet]() that you want to update.")
	waf_updateByteMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	waf_updateByteMatchSetCmd.Flags().String("updates", "", "An array of `ByteMatchSetUpdate` objects that you want to insert into or delete from a [ByteMatchSet]().")
	waf_updateByteMatchSetCmd.MarkFlagRequired("byte-match-set-id")
	waf_updateByteMatchSetCmd.MarkFlagRequired("change-token")
	waf_updateByteMatchSetCmd.MarkFlagRequired("updates")
	wafCmd.AddCommand(waf_updateByteMatchSetCmd)
}
