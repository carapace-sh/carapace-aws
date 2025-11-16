package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createLoadBalancerCmd = &cobra.Command{
	Use:   "create-load-balancer",
	Short: "Creates a Lightsail load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createLoadBalancerCmd).Standalone()

	lightsail_createLoadBalancerCmd.Flags().String("certificate-alternative-names", "", "The optional alternative domains and subdomains to use with your SSL/TLS certificate (`www.example.com`, `example.com`, `m.example.com`, `blog.example.com`).")
	lightsail_createLoadBalancerCmd.Flags().String("certificate-domain-name", "", "The domain name with which your certificate is associated (`example.com`).")
	lightsail_createLoadBalancerCmd.Flags().String("certificate-name", "", "The name of the SSL/TLS certificate.")
	lightsail_createLoadBalancerCmd.Flags().String("health-check-path", "", "The path you provided to perform the load balancer health check.")
	lightsail_createLoadBalancerCmd.Flags().String("instance-port", "", "The instance port where you're creating your load balancer.")
	lightsail_createLoadBalancerCmd.Flags().String("ip-address-type", "", "The IP address type for the load balancer.")
	lightsail_createLoadBalancerCmd.Flags().String("load-balancer-name", "", "The name of your load balancer.")
	lightsail_createLoadBalancerCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
	lightsail_createLoadBalancerCmd.Flags().String("tls-policy-name", "", "The name of the TLS policy to apply to the load balancer.")
	lightsail_createLoadBalancerCmd.MarkFlagRequired("instance-port")
	lightsail_createLoadBalancerCmd.MarkFlagRequired("load-balancer-name")
	lightsailCmd.AddCommand(lightsail_createLoadBalancerCmd)
}
