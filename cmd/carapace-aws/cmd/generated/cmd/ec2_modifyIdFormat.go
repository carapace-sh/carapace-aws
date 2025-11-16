package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyIdFormatCmd = &cobra.Command{
	Use:   "modify-id-format",
	Short: "Modifies the ID format for the specified resource on a per-Region basis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyIdFormatCmd).Standalone()

	ec2_modifyIdFormatCmd.Flags().Bool("no-use-long-ids", false, "Indicate whether the resource should use longer IDs (17-character IDs).")
	ec2_modifyIdFormatCmd.Flags().String("resource", "", "The type of resource: `bundle` | `conversion-task` | `customer-gateway` | `dhcp-options` | `elastic-ip-allocation` | `elastic-ip-association` | `export-task` | `flow-log` | `image` | `import-task` | `internet-gateway` | `network-acl` | `network-acl-association` | `network-interface` | `network-interface-attachment` | `prefix-list` | `route-table` | `route-table-association` | `security-group` | `subnet` | `subnet-cidr-block-association` | `vpc` | `vpc-cidr-block-association` | `vpc-endpoint` | `vpc-peering-connection` | `vpn-connection` | `vpn-gateway`.")
	ec2_modifyIdFormatCmd.Flags().Bool("use-long-ids", false, "Indicate whether the resource should use longer IDs (17-character IDs).")
	ec2_modifyIdFormatCmd.Flag("no-use-long-ids").Hidden = true
	ec2_modifyIdFormatCmd.MarkFlagRequired("no-use-long-ids")
	ec2_modifyIdFormatCmd.MarkFlagRequired("resource")
	ec2_modifyIdFormatCmd.MarkFlagRequired("use-long-ids")
	ec2Cmd.AddCommand(ec2_modifyIdFormatCmd)
}
