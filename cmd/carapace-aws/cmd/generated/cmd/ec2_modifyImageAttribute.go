package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyImageAttributeCmd = &cobra.Command{
	Use:   "modify-image-attribute",
	Short: "Modifies the specified attribute of the specified AMI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyImageAttributeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_modifyImageAttributeCmd).Standalone()

		ec2_modifyImageAttributeCmd.Flags().String("attribute", "", "The name of the attribute to modify.")
		ec2_modifyImageAttributeCmd.Flags().String("description", "", "A new description for the AMI.")
		ec2_modifyImageAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyImageAttributeCmd.Flags().String("image-id", "", "The ID of the AMI.")
		ec2_modifyImageAttributeCmd.Flags().String("imds-support", "", "Set to `v2.0` to indicate that IMDSv2 is specified in the AMI.")
		ec2_modifyImageAttributeCmd.Flags().String("launch-permission", "", "A new launch permission for the AMI.")
		ec2_modifyImageAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_modifyImageAttributeCmd.Flags().String("operation-type", "", "The operation type.")
		ec2_modifyImageAttributeCmd.Flags().String("organization-arns", "", "The Amazon Resource Name (ARN) of an organization.")
		ec2_modifyImageAttributeCmd.Flags().String("organizational-unit-arns", "", "The Amazon Resource Name (ARN) of an organizational unit (OU).")
		ec2_modifyImageAttributeCmd.Flags().String("product-codes", "", "Not supported.")
		ec2_modifyImageAttributeCmd.Flags().String("user-groups", "", "The user groups.")
		ec2_modifyImageAttributeCmd.Flags().String("user-ids", "", "The Amazon Web Services account IDs.")
		ec2_modifyImageAttributeCmd.Flags().String("value", "", "The value of the attribute being modified.")
		ec2_modifyImageAttributeCmd.MarkFlagRequired("image-id")
		ec2_modifyImageAttributeCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_modifyImageAttributeCmd)
}
