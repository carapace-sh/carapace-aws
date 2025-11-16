package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbv2_modifyListenerCmd = &cobra.Command{
	Use:   "modify-listener",
	Short: "Replaces the specified properties of the specified listener.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbv2_modifyListenerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbv2_modifyListenerCmd).Standalone()

		elbv2_modifyListenerCmd.Flags().String("alpn-policy", "", "\\[TLS listeners] The name of the Application-Layer Protocol Negotiation (ALPN) policy.")
		elbv2_modifyListenerCmd.Flags().String("certificates", "", "\\[HTTPS and TLS listeners] The default certificate for the listener.")
		elbv2_modifyListenerCmd.Flags().String("default-actions", "", "The actions for the default rule.")
		elbv2_modifyListenerCmd.Flags().String("listener-arn", "", "The Amazon Resource Name (ARN) of the listener.")
		elbv2_modifyListenerCmd.Flags().String("mutual-authentication", "", "\\[HTTPS listeners] The mutual authentication configuration information.")
		elbv2_modifyListenerCmd.Flags().String("port", "", "The port for connections from clients to the load balancer.")
		elbv2_modifyListenerCmd.Flags().String("protocol", "", "The protocol for connections from clients to the load balancer.")
		elbv2_modifyListenerCmd.Flags().String("ssl-policy", "", "\\[HTTPS and TLS listeners] The security policy that defines which protocols and ciphers are supported.")
		elbv2_modifyListenerCmd.MarkFlagRequired("listener-arn")
	})
	elbv2Cmd.AddCommand(elbv2_modifyListenerCmd)
}
