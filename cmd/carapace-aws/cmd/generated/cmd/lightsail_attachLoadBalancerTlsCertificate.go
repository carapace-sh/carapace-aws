package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_attachLoadBalancerTlsCertificateCmd = &cobra.Command{
	Use:   "attach-load-balancer-tls-certificate",
	Short: "Attaches a Transport Layer Security (TLS) certificate to your load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_attachLoadBalancerTlsCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_attachLoadBalancerTlsCertificateCmd).Standalone()

		lightsail_attachLoadBalancerTlsCertificateCmd.Flags().String("certificate-name", "", "The name of your SSL/TLS certificate.")
		lightsail_attachLoadBalancerTlsCertificateCmd.Flags().String("load-balancer-name", "", "The name of the load balancer to which you want to associate the SSL/TLS certificate.")
		lightsail_attachLoadBalancerTlsCertificateCmd.MarkFlagRequired("certificate-name")
		lightsail_attachLoadBalancerTlsCertificateCmd.MarkFlagRequired("load-balancer-name")
	})
	lightsailCmd.AddCommand(lightsail_attachLoadBalancerTlsCertificateCmd)
}
