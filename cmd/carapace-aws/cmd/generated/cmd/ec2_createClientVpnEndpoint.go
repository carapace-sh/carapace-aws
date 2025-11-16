package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createClientVpnEndpointCmd = &cobra.Command{
	Use:   "create-client-vpn-endpoint",
	Short: "Creates a Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createClientVpnEndpointCmd).Standalone()

	ec2_createClientVpnEndpointCmd.Flags().String("authentication-options", "", "Information about the authentication method to be used to authenticate clients.")
	ec2_createClientVpnEndpointCmd.Flags().String("client-cidr-block", "", "The IPv4 address range, in CIDR notation, from which to assign client IP addresses.")
	ec2_createClientVpnEndpointCmd.Flags().String("client-connect-options", "", "The options for managing connection authorization for new client connections.")
	ec2_createClientVpnEndpointCmd.Flags().String("client-login-banner-options", "", "Options for enabling a customizable text banner that will be displayed on Amazon Web Services provided clients when a VPN session is established.")
	ec2_createClientVpnEndpointCmd.Flags().String("client-route-enforcement-options", "", "Client route enforcement is a feature of the Client VPN service that helps enforce administrator defined routes on devices connected through the VPN.")
	ec2_createClientVpnEndpointCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	ec2_createClientVpnEndpointCmd.Flags().String("connection-log-options", "", "Information about the client connection logging options.")
	ec2_createClientVpnEndpointCmd.Flags().String("description", "", "A brief description of the Client VPN endpoint.")
	ec2_createClientVpnEndpointCmd.Flags().Bool("disconnect-on-session-timeout", false, "Indicates whether the client VPN session is disconnected after the maximum timeout specified in `SessionTimeoutHours` is reached.")
	ec2_createClientVpnEndpointCmd.Flags().String("dns-servers", "", "Information about the DNS servers to be used for DNS resolution.")
	ec2_createClientVpnEndpointCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createClientVpnEndpointCmd.Flags().String("endpoint-ip-address-type", "", "The IP address type for the Client VPN endpoint.")
	ec2_createClientVpnEndpointCmd.Flags().Bool("no-disconnect-on-session-timeout", false, "Indicates whether the client VPN session is disconnected after the maximum timeout specified in `SessionTimeoutHours` is reached.")
	ec2_createClientVpnEndpointCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createClientVpnEndpointCmd.Flags().Bool("no-split-tunnel", false, "Indicates whether split-tunnel is enabled on the Client VPN endpoint.")
	ec2_createClientVpnEndpointCmd.Flags().String("security-group-ids", "", "The IDs of one or more security groups to apply to the target network.")
	ec2_createClientVpnEndpointCmd.Flags().String("self-service-portal", "", "Specify whether to enable the self-service portal for the Client VPN endpoint.")
	ec2_createClientVpnEndpointCmd.Flags().String("server-certificate-arn", "", "The ARN of the server certificate.")
	ec2_createClientVpnEndpointCmd.Flags().String("session-timeout-hours", "", "The maximum VPN session duration time in hours.")
	ec2_createClientVpnEndpointCmd.Flags().Bool("split-tunnel", false, "Indicates whether split-tunnel is enabled on the Client VPN endpoint.")
	ec2_createClientVpnEndpointCmd.Flags().String("tag-specifications", "", "The tags to apply to the Client VPN endpoint during creation.")
	ec2_createClientVpnEndpointCmd.Flags().String("traffic-ip-address-type", "", "The IP address type for traffic within the Client VPN tunnel.")
	ec2_createClientVpnEndpointCmd.Flags().String("transport-protocol", "", "The transport protocol to be used by the VPN session.")
	ec2_createClientVpnEndpointCmd.Flags().String("vpc-id", "", "The ID of the VPC to associate with the Client VPN endpoint.")
	ec2_createClientVpnEndpointCmd.Flags().String("vpn-port", "", "The port number to assign to the Client VPN endpoint for TCP and UDP traffic.")
	ec2_createClientVpnEndpointCmd.MarkFlagRequired("authentication-options")
	ec2_createClientVpnEndpointCmd.MarkFlagRequired("connection-log-options")
	ec2_createClientVpnEndpointCmd.Flag("no-disconnect-on-session-timeout").Hidden = true
	ec2_createClientVpnEndpointCmd.Flag("no-dry-run").Hidden = true
	ec2_createClientVpnEndpointCmd.Flag("no-split-tunnel").Hidden = true
	ec2_createClientVpnEndpointCmd.MarkFlagRequired("server-certificate-arn")
	ec2Cmd.AddCommand(ec2_createClientVpnEndpointCmd)
}
