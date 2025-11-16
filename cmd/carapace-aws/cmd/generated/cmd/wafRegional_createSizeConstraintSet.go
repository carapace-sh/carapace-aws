package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_createSizeConstraintSetCmd = &cobra.Command{
	Use:   "create-size-constraint-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_createSizeConstraintSetCmd).Standalone()

	wafRegional_createSizeConstraintSetCmd.Flags().String("change-token", "", "The value returned by the most recent call to [GetChangeToken]().")
	wafRegional_createSizeConstraintSetCmd.Flags().String("name", "", "A friendly name or description of the [SizeConstraintSet]().")
	wafRegional_createSizeConstraintSetCmd.MarkFlagRequired("change-token")
	wafRegional_createSizeConstraintSetCmd.MarkFlagRequired("name")
	wafRegionalCmd.AddCommand(wafRegional_createSizeConstraintSetCmd)
}
