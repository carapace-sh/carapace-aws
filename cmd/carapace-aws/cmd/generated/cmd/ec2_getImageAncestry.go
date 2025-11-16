package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getImageAncestryCmd = &cobra.Command{
	Use:   "get-image-ancestry",
	Short: "Retrieves the ancestry chain of the specified AMI, tracing its lineage back to the root AMI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getImageAncestryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getImageAncestryCmd).Standalone()

		ec2_getImageAncestryCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getImageAncestryCmd.Flags().String("image-id", "", "The ID of the AMI whose ancestry you want to trace.")
		ec2_getImageAncestryCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getImageAncestryCmd.MarkFlagRequired("image-id")
		ec2_getImageAncestryCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_getImageAncestryCmd)
}
