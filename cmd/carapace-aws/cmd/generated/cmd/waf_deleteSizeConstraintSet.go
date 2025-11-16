package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_deleteSizeConstraintSetCmd = &cobra.Command{
	Use:   "delete-size-constraint-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_deleteSizeConstraintSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_deleteSizeConstraintSetCmd).Standalone()

		waf_deleteSizeConstraintSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
		waf_deleteSizeConstraintSetCmd.Flags().String("size-constraint-set-id", "", "The `SizeConstraintSetId` of the [SizeConstraintSet]() that you want to delete.")
		waf_deleteSizeConstraintSetCmd.MarkFlagRequired("change-token")
		waf_deleteSizeConstraintSetCmd.MarkFlagRequired("size-constraint-set-id")
	})
	wafCmd.AddCommand(waf_deleteSizeConstraintSetCmd)
}
