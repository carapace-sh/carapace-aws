package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createTagsCmd = &cobra.Command{
	Use:   "create-tags",
	Short: "Adds or overwrites only the specified tags for the specified Amazon EC2 resource or resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createTagsCmd).Standalone()

		ec2_createTagsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTagsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_createTagsCmd.Flags().String("resources", "", "The IDs of the resources, separated by spaces.")
		ec2_createTagsCmd.Flags().String("tags", "", "The tags.")
		ec2_createTagsCmd.Flag("no-dry-run").Hidden = true
		ec2_createTagsCmd.MarkFlagRequired("resources")
		ec2_createTagsCmd.MarkFlagRequired("tags")
	})
	ec2Cmd.AddCommand(ec2_createTagsCmd)
}
