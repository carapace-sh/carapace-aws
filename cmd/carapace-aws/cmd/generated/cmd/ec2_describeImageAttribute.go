package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeImageAttributeCmd = &cobra.Command{
	Use:   "describe-image-attribute",
	Short: "Describes the specified attribute of the specified AMI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeImageAttributeCmd).Standalone()

	ec2_describeImageAttributeCmd.Flags().String("attribute", "", "The AMI attribute.")
	ec2_describeImageAttributeCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeImageAttributeCmd.Flags().String("image-id", "", "The ID of the AMI.")
	ec2_describeImageAttributeCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeImageAttributeCmd.MarkFlagRequired("attribute")
	ec2_describeImageAttributeCmd.MarkFlagRequired("image-id")
	ec2_describeImageAttributeCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeImageAttributeCmd)
}
