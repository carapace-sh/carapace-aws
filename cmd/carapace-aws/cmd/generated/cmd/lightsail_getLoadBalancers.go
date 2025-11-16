package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getLoadBalancersCmd = &cobra.Command{
	Use:   "get-load-balancers",
	Short: "Returns information about all load balancers in an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getLoadBalancersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getLoadBalancersCmd).Standalone()

		lightsail_getLoadBalancersCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	})
	lightsailCmd.AddCommand(lightsail_getLoadBalancersCmd)
}
