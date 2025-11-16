package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waf_listIpsetsCmd = &cobra.Command{
	Use:   "list-ipsets",
	Short: "This is **AWS WAF Classic** documentation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waf_listIpsetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(waf_listIpsetsCmd).Standalone()

		waf_listIpsetsCmd.Flags().String("limit", "", "Specifies the number of `IPSet` objects that you want AWS WAF to return for this request.")
		waf_listIpsetsCmd.Flags().String("next-marker", "", "AWS WAF returns a `NextMarker` value in the response that allows you to list another group of `IPSets`.")
	})
	wafCmd.AddCommand(waf_listIpsetsCmd)
}
