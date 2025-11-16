package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getReusableDelegationSetCmd = &cobra.Command{
	Use:   "get-reusable-delegation-set",
	Short: "Retrieves information about a specified reusable delegation set, including the four name servers that are assigned to the delegation set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getReusableDelegationSetCmd).Standalone()

	route53_getReusableDelegationSetCmd.Flags().String("id", "", "The ID of the reusable delegation set that you want to get a list of name servers for.")
	route53_getReusableDelegationSetCmd.MarkFlagRequired("id")
	route53Cmd.AddCommand(route53_getReusableDelegationSetCmd)
}
