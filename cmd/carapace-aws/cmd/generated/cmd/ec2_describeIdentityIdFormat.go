package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeIdentityIdFormatCmd = &cobra.Command{
	Use:   "describe-identity-id-format",
	Short: "Describes the ID format settings for resources for the specified IAM user, IAM role, or root user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeIdentityIdFormatCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeIdentityIdFormatCmd).Standalone()

		ec2_describeIdentityIdFormatCmd.Flags().String("principal-arn", "", "The ARN of the principal, which can be an IAM role, IAM user, or the root user.")
		ec2_describeIdentityIdFormatCmd.Flags().String("resource", "", "The type of resource: `bundle` | `conversion-task` | `customer-gateway` | `dhcp-options` | `elastic-ip-allocation` | `elastic-ip-association` | `export-task` | `flow-log` | `image` | `import-task` | `instance` | `internet-gateway` | `network-acl` | `network-acl-association` | `network-interface` | `network-interface-attachment` | `prefix-list` | `reservation` | `route-table` | `route-table-association` | `security-group` | `snapshot` | `subnet` | `subnet-cidr-block-association` | `volume` | `vpc` | `vpc-cidr-block-association` | `vpc-endpoint` | `vpc-peering-connection` | `vpn-connection` | `vpn-gateway`")
		ec2_describeIdentityIdFormatCmd.MarkFlagRequired("principal-arn")
	})
	ec2Cmd.AddCommand(ec2_describeIdentityIdFormatCmd)
}
