package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeSecurityGroupReferencesCmd = &cobra.Command{
	Use:   "describe-security-group-references",
	Short: "Describes the VPCs on the other side of a VPC peering or Transit Gateway connection that are referencing the security groups you've specified in this request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeSecurityGroupReferencesCmd).Standalone()

	ec2_describeSecurityGroupReferencesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeSecurityGroupReferencesCmd.Flags().String("group-id", "", "The IDs of the security groups in your account.")
	ec2_describeSecurityGroupReferencesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeSecurityGroupReferencesCmd.MarkFlagRequired("group-id")
	ec2_describeSecurityGroupReferencesCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeSecurityGroupReferencesCmd)
}
