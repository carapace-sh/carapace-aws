package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeVpcAttributeCmd = &cobra.Command{
	Use:   "describe-vpc-attribute",
	Short: "Describes the specified attribute of the specified VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeVpcAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeVpcAttributeCmd).Standalone()

		ec2_describeVpcAttributeCmd.Flags().String("attribute", "", "The VPC attribute.")
		ec2_describeVpcAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeVpcAttributeCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
		ec2_describeVpcAttributeCmd.MarkFlagRequired("attribute")
		ec2_describeVpcAttributeCmd.Flag("no-dry-run").Hidden = true
		ec2_describeVpcAttributeCmd.MarkFlagRequired("vpc-id")
	})
	ec2Cmd.AddCommand(ec2_describeVpcAttributeCmd)
}
