package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteLoadBalancerTlsCertificateCmd = &cobra.Command{
	Use:   "delete-load-balancer-tls-certificate",
	Short: "Deletes an SSL/TLS certificate associated with a Lightsail load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteLoadBalancerTlsCertificateCmd).Standalone()

	lightsail_deleteLoadBalancerTlsCertificateCmd.Flags().String("certificate-name", "", "The SSL/TLS certificate name.")
	lightsail_deleteLoadBalancerTlsCertificateCmd.Flags().String("force", "", "When `true`, forces the deletion of an SSL/TLS certificate.")
	lightsail_deleteLoadBalancerTlsCertificateCmd.Flags().String("load-balancer-name", "", "The load balancer name.")
	lightsail_deleteLoadBalancerTlsCertificateCmd.MarkFlagRequired("certificate-name")
	lightsail_deleteLoadBalancerTlsCertificateCmd.MarkFlagRequired("load-balancer-name")
	lightsailCmd.AddCommand(lightsail_deleteLoadBalancerTlsCertificateCmd)
}
