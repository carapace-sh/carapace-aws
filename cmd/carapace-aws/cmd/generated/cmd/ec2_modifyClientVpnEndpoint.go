package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyClientVpnEndpointCmd = &cobra.Command{
	Use:   "modify-client-vpn-endpoint",
	Short: "Modifies the specified Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyClientVpnEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyClientVpnEndpointCmd).Standalone()

		ec2_modifyClientVpnEndpointCmd.Flags().String("client-connect-options", "", "The options for managing connection authorization for new client connections.")
		ec2_modifyClientVpnEndpointCmd.Flags().String("client-login-banner-options", "", "Options for enabling a customizable text banner that will be displayed on Amazon Web Services provided clients when a VPN session is established.")
		ec2_modifyClientVpnEndpointCmd.Flags().String("client-route-enforcement-options", "", "Client route enforcement is a feature of the Client VPN service that helps enforce administrator defined routes on devices connected through the VPN.")
		ec2_modifyClientVpnEndpointCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint to modify.")
		ec2_modifyClientVpnEndpointCmd.Flags().String("connection-log-options", "", "Information about the client connection logging options.")
		ec2_modifyClientVpnEndpointCmd.Flags().String("description", "", "A brief description of the Client VPN endpoint.")
		ec2_modifyClientVpnEndpointCmd.Flags().Bool("disconnect-on-session-timeout", false, "Indicates whether the client VPN session is disconnected after the maximum timeout specified in `sessionTimeoutHours` is reached.")
		ec2_modifyClientVpnEndpointCmd.Flags().String("dns-servers", "", "Information about the DNS servers to be used by Client VPN connections.")
		ec2_modifyClientVpnEndpointCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyClientVpnEndpointCmd.Flags().Bool("no-disconnect-on-session-timeout", false, "Indicates whether the client VPN session is disconnected after the maximum timeout specified in `sessionTimeoutHours` is reached.")
		ec2_modifyClientVpnEndpointCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyClientVpnEndpointCmd.Flags().Bool("no-split-tunnel", false, "Indicates whether the VPN is split-tunnel.")
		ec2_modifyClientVpnEndpointCmd.Flags().String("security-group-ids", "", "The IDs of one or more security groups to apply to the target network.")
		ec2_modifyClientVpnEndpointCmd.Flags().String("self-service-portal", "", "Specify whether to enable the self-service portal for the Client VPN endpoint.")
		ec2_modifyClientVpnEndpointCmd.Flags().String("server-certificate-arn", "", "The ARN of the server certificate to be used.")
		ec2_modifyClientVpnEndpointCmd.Flags().String("session-timeout-hours", "", "The maximum VPN session duration time in hours.")
		ec2_modifyClientVpnEndpointCmd.Flags().Bool("split-tunnel", false, "Indicates whether the VPN is split-tunnel.")
		ec2_modifyClientVpnEndpointCmd.Flags().String("vpc-id", "", "The ID of the VPC to associate with the Client VPN endpoint.")
		ec2_modifyClientVpnEndpointCmd.Flags().String("vpn-port", "", "The port number to assign to the Client VPN endpoint for TCP and UDP traffic.")
		ec2_modifyClientVpnEndpointCmd.MarkFlagRequired("client-vpn-endpoint-id")
		ec2_modifyClientVpnEndpointCmd.Flag("no-disconnect-on-session-timeout").Hidden = true
		ec2_modifyClientVpnEndpointCmd.Flag("no-dry-run").Hidden = true
		ec2_modifyClientVpnEndpointCmd.Flag("no-split-tunnel").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyClientVpnEndpointCmd)
}
