package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeClientVpnAuthorizationRulesCmd = &cobra.Command{
	Use:   "describe-client-vpn-authorization-rules",
	Short: "Describes the authorization rules for a specified Client VPN endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeClientVpnAuthorizationRulesCmd).Standalone()

	ec2_describeClientVpnAuthorizationRulesCmd.Flags().String("client-vpn-endpoint-id", "", "The ID of the Client VPN endpoint.")
	ec2_describeClientVpnAuthorizationRulesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeClientVpnAuthorizationRulesCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeClientVpnAuthorizationRulesCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
	ec2_describeClientVpnAuthorizationRulesCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	ec2_describeClientVpnAuthorizationRulesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeClientVpnAuthorizationRulesCmd.MarkFlagRequired("client-vpn-endpoint-id")
	ec2_describeClientVpnAuthorizationRulesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeClientVpnAuthorizationRulesCmd)
}
