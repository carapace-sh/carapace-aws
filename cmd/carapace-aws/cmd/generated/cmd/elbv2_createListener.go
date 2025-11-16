package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_createListenerCmd = &cobra.Command{
	Use:   "create-listener",
	Short: "Creates a listener for the specified Application Load Balancer, Network Load Balancer, or Gateway Load Balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_createListenerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_createListenerCmd).Standalone()

		elbv2_createListenerCmd.Flags().String("alpn-policy", "", "\\[TLS listeners] The name of the Application-Layer Protocol Negotiation (ALPN) policy.")
		elbv2_createListenerCmd.Flags().String("certificates", "", "\\[HTTPS and TLS listeners] The default certificate for the listener.")
		elbv2_createListenerCmd.Flags().String("default-actions", "", "The actions for the default rule.")
		elbv2_createListenerCmd.Flags().String("load-balancer-arn", "", "The Amazon Resource Name (ARN) of the load balancer.")
		elbv2_createListenerCmd.Flags().String("mutual-authentication", "", "\\[HTTPS listeners] The mutual authentication configuration information.")
		elbv2_createListenerCmd.Flags().String("port", "", "The port on which the load balancer is listening.")
		elbv2_createListenerCmd.Flags().String("protocol", "", "The protocol for connections from clients to the load balancer.")
		elbv2_createListenerCmd.Flags().String("ssl-policy", "", "\\[HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.")
		elbv2_createListenerCmd.Flags().String("tags", "", "The tags to assign to the listener.")
		elbv2_createListenerCmd.MarkFlagRequired("default-actions")
		elbv2_createListenerCmd.MarkFlagRequired("load-balancer-arn")
	})
	elbv2Cmd.AddCommand(elbv2_createListenerCmd)
}
