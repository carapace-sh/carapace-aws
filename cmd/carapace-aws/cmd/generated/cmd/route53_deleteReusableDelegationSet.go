package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_deleteReusableDelegationSetCmd = &cobra.Command{
	Use:   "delete-reusable-delegation-set",
	Short: "Deletes a reusable delegation set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_deleteReusableDelegationSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_deleteReusableDelegationSetCmd).Standalone()

		route53_deleteReusableDelegationSetCmd.Flags().String("id", "", "The ID of the reusable delegation set that you want to delete.")
		route53_deleteReusableDelegationSetCmd.MarkFlagRequired("id")
	})
	route53Cmd.AddCommand(route53_deleteReusableDelegationSetCmd)
}
