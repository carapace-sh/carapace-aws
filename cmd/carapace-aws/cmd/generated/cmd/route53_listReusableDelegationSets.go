package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listReusableDelegationSetsCmd = &cobra.Command{
	Use:   "list-reusable-delegation-sets",
	Short: "Retrieves a list of the reusable delegation sets that are associated with the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listReusableDelegationSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_listReusableDelegationSetsCmd).Standalone()

		route53_listReusableDelegationSetsCmd.Flags().String("marker", "", "If the value of `IsTruncated` in the previous response was `true`, you have more reusable delegation sets.")
		route53_listReusableDelegationSetsCmd.Flags().String("max-items", "", "The number of reusable delegation sets that you want Amazon Route 53 to return in the response to this request.")
	})
	route53Cmd.AddCommand(route53_listReusableDelegationSetsCmd)
}
