package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53Cmd = &cobra.Command{
	Use:   "route53",
	Short: "Amazon Route\u00a053 is a highly available and scalable Domain Name System (DNS) web service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53Cmd).Standalone()

	})
	rootCmd.AddCommand(route53Cmd)
}
