package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeInstanceAttributeCmd = &cobra.Command{
	Use:   "describe-instance-attribute",
	Short: "Describes the specified attribute of the specified instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeInstanceAttributeCmd).Standalone()

	ec2_describeInstanceAttributeCmd.Flags().String("attribute", "", "The instance attribute.")
	ec2_describeInstanceAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_describeInstanceAttributeCmd.Flags().String("instance-id", "", "The ID of the instance.")
	ec2_describeInstanceAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the operation, without actually making the request, and provides an error response.")
	ec2_describeInstanceAttributeCmd.MarkFlagRequired("attribute")
	ec2_describeInstanceAttributeCmd.MarkFlagRequired("instance-id")
	ec2_describeInstanceAttributeCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeInstanceAttributeCmd)
}
