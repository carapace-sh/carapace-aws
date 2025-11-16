package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafRegional_listSizeConstraintSetsCmd = &cobra.Command{
	Use:   "list-size-constraint-sets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafRegional_listSizeConstraintSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafRegional_listSizeConstraintSetsCmd).Standalone()

		wafRegional_listSizeConstraintSetsCmd.Flags().String("limit", "", "Specifies the number of `SizeConstraintSet` objects that you want AWS WAF to return for this request.")
		wafRegional_listSizeConstraintSetsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `SizeConstraintSets` than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `SizeConstraintSets`.")
	})
	wafRegionalCmd.AddCommand(wafRegional_listSizeConstraintSetsCmd)
}
