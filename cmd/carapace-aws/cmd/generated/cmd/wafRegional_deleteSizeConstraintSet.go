package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_deleteSizeConstraintSetCmd = &cobra.Command{
	Use:   "delete-size-constraint-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_deleteSizeConstraintSetCmd).Standalone()

	wafRegional_deleteSizeConstraintSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_deleteSizeConstraintSetCmd.Flags().String("size-constraint-set-id", "", "The `SizeConstraintSetId` of the [SizeConstraintSet]() that you want to delete.")
	wafRegional_deleteSizeConstraintSetCmd.MarkFlagRequired("change-token")
	wafRegional_deleteSizeConstraintSetCmd.MarkFlagRequired("size-constraint-set-id")
	wafRegionalCmd.AddCommand(wafRegional_deleteSizeConstraintSetCmd)
}
