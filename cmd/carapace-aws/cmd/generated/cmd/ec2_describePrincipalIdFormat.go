package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describePrincipalIdFormatCmd = &cobra.Command{
	Use:   "describe-principal-id-format",
	Short: "Describes the ID format settings for the root user and all IAM roles and IAM users that have explicitly specified a longer ID (17-character ID) preference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describePrincipalIdFormatCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describePrincipalIdFormatCmd).Standalone()

		ec2_describePrincipalIdFormatCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describePrincipalIdFormatCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		ec2_describePrincipalIdFormatCmd.Flags().String("next-token", "", "The token to request the next page of results.")
		ec2_describePrincipalIdFormatCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describePrincipalIdFormatCmd.Flags().String("resources", "", "The type of resource: `bundle` | `conversion-task` | `customer-gateway` | `dhcp-options` | `elastic-ip-allocation` | `elastic-ip-association` | `export-task` | `flow-log` | `image` | `import-task` | `instance` | `internet-gateway` | `network-acl` | `network-acl-association` | `network-interface` | `network-interface-attachment` | `prefix-list` | `reservation` | `route-table` | `route-table-association` | `security-group` | `snapshot` | `subnet` | `subnet-cidr-block-association` | `volume` | `vpc` | `vpc-cidr-block-association` | `vpc-endpoint` | `vpc-peering-connection` | `vpn-connection` | `vpn-gateway`")
		ec2_describePrincipalIdFormatCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describePrincipalIdFormatCmd)
}
