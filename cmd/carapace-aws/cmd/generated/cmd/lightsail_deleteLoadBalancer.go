package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteLoadBalancerCmd = &cobra.Command{
	Use:   "delete-load-balancer",
	Short: "Deletes a Lightsail load balancer and all its associated SSL/TLS certificates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteLoadBalancerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_deleteLoadBalancerCmd).Standalone()

		lightsail_deleteLoadBalancerCmd.Flags().String("load-balancer-name", "", "The name of the load balancer you want to delete.")
		lightsail_deleteLoadBalancerCmd.MarkFlagRequired("load-balancer-name")
	})
	lightsailCmd.AddCommand(lightsail_deleteLoadBalancerCmd)
}
