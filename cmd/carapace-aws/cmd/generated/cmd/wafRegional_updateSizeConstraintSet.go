package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_updateSizeConstraintSetCmd = &cobra.Command{
	Use:   "update-size-constraint-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_updateSizeConstraintSetCmd).Standalone()

	wafRegional_updateSizeConstraintSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_updateSizeConstraintSetCmd.Flags().String("size-constraint-set-id", "", "The `SizeConstraintSetId` of the [SizeConstraintSet]() that you want to update.")
	wafRegional_updateSizeConstraintSetCmd.Flags().String("updates", "", "An array of `SizeConstraintSetUpdate` objects that you want to insert into or delete from a [SizeConstraintSet]().")
	wafRegional_updateSizeConstraintSetCmd.MarkFlagRequired("change-token")
	wafRegional_updateSizeConstraintSetCmd.MarkFlagRequired("size-constraint-set-id")
	wafRegional_updateSizeConstraintSetCmd.MarkFlagRequired("updates")
	wafRegionalCmd.AddCommand(wafRegional_updateSizeConstraintSetCmd)
}
