package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyIdentityIdFormatCmd = &cobra.Command{
	Use:   "modify-identity-id-format",
	Short: "Modifies the ID format of a resource for a specified IAM user, IAM role, or the root user for an account; or all IAM users, IAM roles, and the root user for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyIdentityIdFormatCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyIdentityIdFormatCmd).Standalone()

		ec2_modifyIdentityIdFormatCmd.Flags().Bool("no-use-long-ids", false, "Indicates whether the resource should use longer IDs (17-character IDs)")
		ec2_modifyIdentityIdFormatCmd.Flags().String("principal-arn", "", "The ARN of the principal, which can be an IAM user, IAM role, or the root user.")
		ec2_modifyIdentityIdFormatCmd.Flags().String("resource", "", "The type of resource: `bundle` | `conversion-task` | `customer-gateway` | `dhcp-options` | `elastic-ip-allocation` | `elastic-ip-association` | `export-task` | `flow-log` | `image` | `import-task` | `internet-gateway` | `network-acl` | `network-acl-association` | `network-interface` | `network-interface-attachment` | `prefix-list` | `route-table` | `route-table-association` | `security-group` | `subnet` | `subnet-cidr-block-association` | `vpc` | `vpc-cidr-block-association` | `vpc-endpoint` | `vpc-peering-connection` | `vpn-connection` | `vpn-gateway`.")
		ec2_modifyIdentityIdFormatCmd.Flags().Bool("use-long-ids", false, "Indicates whether the resource should use longer IDs (17-character IDs)")
		ec2_modifyIdentityIdFormatCmd.Flag("no-use-long-ids").Hidden = true
		ec2_modifyIdentityIdFormatCmd.MarkFlagRequired("no-use-long-ids")
		ec2_modifyIdentityIdFormatCmd.MarkFlagRequired("principal-arn")
		ec2_modifyIdentityIdFormatCmd.MarkFlagRequired("resource")
		ec2_modifyIdentityIdFormatCmd.MarkFlagRequired("use-long-ids")
	})
	ec2Cmd.AddCommand(ec2_modifyIdentityIdFormatCmd)
}
