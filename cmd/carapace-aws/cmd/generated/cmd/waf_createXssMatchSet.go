package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_createXssMatchSetCmd = &cobra.Command{
	Use:   "create-xss-match-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_createXssMatchSetCmd).Standalone()

	waf_createXssMatchSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	waf_createXssMatchSetCmd.Flags().String("name", "", "A friendly name or description for the [XssMatchSet]() that you're creating.")
	waf_createXssMatchSetCmd.MarkFlagRequired("change-token")
	waf_createXssMatchSetCmd.MarkFlagRequired("name")
	wafCmd.AddCommand(waf_createXssMatchSetCmd)
}
