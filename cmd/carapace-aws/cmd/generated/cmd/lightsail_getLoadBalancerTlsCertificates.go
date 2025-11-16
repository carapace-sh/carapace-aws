package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getLoadBalancerTlsCertificatesCmd = &cobra.Command{
	Use:   "get-load-balancer-tls-certificates",
	Short: "Returns information about the TLS certificates that are associated with the specified Lightsail load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getLoadBalancerTlsCertificatesCmd).Standalone()

	lightsail_getLoadBalancerTlsCertificatesCmd.Flags().String("load-balancer-name", "", "The name of the load balancer you associated with your SSL/TLS certificate.")
	lightsail_getLoadBalancerTlsCertificatesCmd.MarkFlagRequired("load-balancer-name")
	lightsailCmd.AddCommand(lightsail_getLoadBalancerTlsCertificatesCmd)
}
