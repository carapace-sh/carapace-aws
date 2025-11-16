package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_listSizeConstraintSetsCmd = &cobra.Command{
	Use:   "list-size-constraint-sets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_listSizeConstraintSetsCmd).Standalone()

	waf_listSizeConstraintSetsCmd.Flags().String("limit", "", "Specifies the number of `SizeConstraintSet` objects that you want AWS WAF to return for this request.")
	waf_listSizeConstraintSetsCmd.Flags().String("next-marker", "", "If you specify a value for `Limit` and you have more `SizeConstraintSets` than the value of `Limit`, AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `SizeConstraintSets`.")
	wafCmd.AddCommand(waf_listSizeConstraintSetsCmd)
}
