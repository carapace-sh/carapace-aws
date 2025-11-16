package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_createByteMatchSetCmd = &cobra.Command{
	Use:   "create-byte-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_createByteMatchSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_createByteMatchSetCmd).Standalone()

		waf_createByteMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_createByteMatchSetCmd.Flags().String("name", "", "A friendly name or description of the [ByteMatchSet]().")
		waf_createByteMatchSetCmd.MarkFlagRequired("change-token")
		waf_createByteMatchSetCmd.MarkFlagRequired("name")
	})
	wafCmd.AddCommand(waf_createByteMatchSetCmd)
}
