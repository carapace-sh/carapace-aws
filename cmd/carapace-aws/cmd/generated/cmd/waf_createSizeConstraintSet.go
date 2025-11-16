package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_createSizeConstraintSetCmd = &cobra.Command{
	Use:   "create-size-constraint-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_createSizeConstraintSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_createSizeConstraintSetCmd).Standalone()

		waf_createSizeConstraintSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_createSizeConstraintSetCmd.Flags().String("name", "", "A friendly name or description of the [SizeConstraintSet]().")
		waf_createSizeConstraintSetCmd.MarkFlagRequired("change-token")
		waf_createSizeConstraintSetCmd.MarkFlagRequired("name")
	})
	wafCmd.AddCommand(waf_createSizeConstraintSetCmd)
}
