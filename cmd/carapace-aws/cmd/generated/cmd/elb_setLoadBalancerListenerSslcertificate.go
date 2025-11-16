package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_setLoadBalancerListenerSslcertificateCmd = &cobra.Command{
	Use:   "set-load-balancer-listener-sslcertificate",
	Short: "Sets the certificate that terminates the specified listener's SSL connections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_setLoadBalancerListenerSslcertificateCmd).Standalone()

	elb_setLoadBalancerListenerSslcertificateCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
	elb_setLoadBalancerListenerSslcertificateCmd.Flags().String("load-balancer-port", "", "The port that uses the specified SSL certificate.")
	elb_setLoadBalancerListenerSslcertificateCmd.Flags().String("sslcertificate-id", "", "The Amazon Resource Name (ARN) of the SSL certificate.")
	elb_setLoadBalancerListenerSslcertificateCmd.MarkFlagRequired("load-balancer-name")
	elb_setLoadBalancerListenerSslcertificateCmd.MarkFlagRequired("load-balancer-port")
	elb_setLoadBalancerListenerSslcertificateCmd.MarkFlagRequired("sslcertificate-id")
	elbCmd.AddCommand(elb_setLoadBalancerListenerSslcertificateCmd)
}
