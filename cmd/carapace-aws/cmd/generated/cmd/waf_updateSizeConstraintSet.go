package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_updateSizeConstraintSetCmd = &cobra.Command{
	Use:   "update-size-constraint-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_updateSizeConstraintSetCmd).Standalone()

	waf_updateSizeConstraintSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	waf_updateSizeConstraintSetCmd.Flags().String("size-constraint-set-id", "", "The `SizeConstraintSetId` of the [SizeConstraintSet]() that you want to update.")
	waf_updateSizeConstraintSetCmd.Flags().String("updates", "", "An array of `SizeConstraintSetUpdate` objects that you want to insert into or delete from a [SizeConstraintSet]().")
	waf_updateSizeConstraintSetCmd.MarkFlagRequired("change-token")
	waf_updateSizeConstraintSetCmd.MarkFlagRequired("size-constraint-set-id")
	waf_updateSizeConstraintSetCmd.MarkFlagRequired("updates")
	wafCmd.AddCommand(waf_updateSizeConstraintSetCmd)
}
