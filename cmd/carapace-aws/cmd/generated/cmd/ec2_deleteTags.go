package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteTagsCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "Deletes the specified set of tags from the specified set of resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteTagsCmd).Standalone()

	ec2_deleteTagsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteTagsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_deleteTagsCmd.Flags().String("resources", "", "The IDs of the resources, separated by spaces.")
	ec2_deleteTagsCmd.Flags().String("tags", "", "The tags to delete.")
	ec2_deleteTagsCmd.Flag("no-dry-run").Hidden = true
	ec2_deleteTagsCmd.MarkFlagRequired("resources")
	ec2Cmd.AddCommand(ec2_deleteTagsCmd)
}
