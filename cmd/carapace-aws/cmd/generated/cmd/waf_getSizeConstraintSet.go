package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_getSizeConstraintSetCmd = &cobra.Command{
	Use:   "get-size-constraint-set",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_getSizeConstraintSetCmd).Standalone()

	waf_getSizeConstraintSetCmd.Flags().String("size-constraint-set-id", "", "The `SizeConstraintSetId` of the [SizeConstraintSet]() that you want to get.")
	waf_getSizeConstraintSetCmd.MarkFlagRequired("size-constraint-set-id")
	wafCmd.AddCommand(waf_getSizeConstraintSetCmd)
}
