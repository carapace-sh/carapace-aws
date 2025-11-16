package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getReusableDelegationSetLimitCmd = &cobra.Command{
	Use:   "get-reusable-delegation-set-limit",
	Short: "Gets the maximum number of hosted zones that you can associate with the specified reusable delegation set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getReusableDelegationSetLimitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_getReusableDelegationSetLimitCmd).Standalone()

		route53_getReusableDelegationSetLimitCmd.Flags().String("delegation-set-id", "", "The ID of the delegation set that you want to get the limit for.")
		route53_getReusableDelegationSetLimitCmd.Flags().String("type", "", "Specify `MAX_ZONES_BY_REUSABLE_DELEGATION_SET` to get the maximum number of hosted zones that you can associate with the specified reusable delegation set.")
		route53_getReusableDelegationSetLimitCmd.MarkFlagRequired("delegation-set-id")
		route53_getReusableDelegationSetLimitCmd.MarkFlagRequired("type")
	})
	route53Cmd.AddCommand(route53_getReusableDelegationSetLimitCmd)
}
