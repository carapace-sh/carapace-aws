package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createLoadBalancerTlsCertificateCmd = &cobra.Command{
	Use:   "create-load-balancer-tls-certificate",
	Short: "Creates an SSL/TLS certificate for an Amazon Lightsail load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createLoadBalancerTlsCertificateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_createLoadBalancerTlsCertificateCmd).Standalone()

		lightsail_createLoadBalancerTlsCertificateCmd.Flags().String("certificate-alternative-names", "", "An array of strings listing alternative domains and subdomains for your SSL/TLS certificate.")
		lightsail_createLoadBalancerTlsCertificateCmd.Flags().String("certificate-domain-name", "", "The domain name (`example.com`) for your SSL/TLS certificate.")
		lightsail_createLoadBalancerTlsCertificateCmd.Flags().String("certificate-name", "", "The SSL/TLS certificate name.")
		lightsail_createLoadBalancerTlsCertificateCmd.Flags().String("load-balancer-name", "", "The load balancer name where you want to create the SSL/TLS certificate.")
		lightsail_createLoadBalancerTlsCertificateCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
		lightsail_createLoadBalancerTlsCertificateCmd.MarkFlagRequired("certificate-domain-name")
		lightsail_createLoadBalancerTlsCertificateCmd.MarkFlagRequired("certificate-name")
		lightsail_createLoadBalancerTlsCertificateCmd.MarkFlagRequired("load-balancer-name")
	})
	lightsailCmd.AddCommand(lightsail_createLoadBalancerTlsCertificateCmd)
}
