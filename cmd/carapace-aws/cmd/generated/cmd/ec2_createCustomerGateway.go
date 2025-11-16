package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createCustomerGatewayCmd = &cobra.Command{
	Use:   "create-customer-gateway",
	Short: "Provides information to Amazon Web Services about your customer gateway device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createCustomerGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createCustomerGatewayCmd).Standalone()

		ec2_createCustomerGatewayCmd.Flags().String("bgp-asn", "", "For customer gateway devices that support BGP, specify the device's ASN.")
		ec2_createCustomerGatewayCmd.Flags().String("bgp-asn-extended", "", "For customer gateway devices that support BGP, specify the device's ASN.")
		ec2_createCustomerGatewayCmd.Flags().String("certificate-arn", "", "The Amazon Resource Name (ARN) for the customer gateway certificate.")
		ec2_createCustomerGatewayCmd.Flags().String("device-name", "", "A name for the customer gateway device.")
		ec2_createCustomerGatewayCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createCustomerGatewayCmd.Flags().String("ip-address", "", "The IP address for the customer gateway device's outside interface.")
		ec2_createCustomerGatewayCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createCustomerGatewayCmd.Flags().String("public-ip", "", "*This member has been deprecated.* The Internet-routable IP address for the customer gateway's outside interface.")
		ec2_createCustomerGatewayCmd.Flags().String("tag-specifications", "", "The tags to apply to the customer gateway.")
		ec2_createCustomerGatewayCmd.Flags().String("type", "", "The type of VPN connection that this customer gateway supports (`ipsec.1`).")
		ec2_createCustomerGatewayCmd.Flag("no-dry-run").Hidden = true
		ec2_createCustomerGatewayCmd.MarkFlagRequired("type")
	})
	ec2Cmd.AddCommand(ec2_createCustomerGatewayCmd)
}
