package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createByteMatchSetCmd = &cobra.Command{
	Use:   "create-byte-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createByteMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_createByteMatchSetCmd).Standalone()

		wafRegional_createByteMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		wafRegional_createByteMatchSetCmd.Flags().String("name", "", "A friendly name or description of the [ByteMatchSet]().")
		wafRegional_createByteMatchSetCmd.MarkFlagRequired("change-token")
		wafRegional_createByteMatchSetCmd.MarkFlagRequired("name")
	})
	wafRegionalCmd.AddCommand(wafRegional_createByteMatchSetCmd)
}
