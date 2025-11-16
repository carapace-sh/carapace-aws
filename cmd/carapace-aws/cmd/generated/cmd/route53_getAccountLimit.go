package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_getAccountLimitCmd = &cobra.Command{
	Use:   "get-account-limit",
	Short: "Gets the specified limit for the current account, for example, the maximum number of health checks that you can create using the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_getAccountLimitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_getAccountLimitCmd).Standalone()

		route53_getAccountLimitCmd.Flags().String("type", "", "The limit that you want to get.")
		route53_getAccountLimitCmd.MarkFlagRequired("type")
	})
	route53Cmd.AddCommand(route53_getAccountLimitCmd)
}
