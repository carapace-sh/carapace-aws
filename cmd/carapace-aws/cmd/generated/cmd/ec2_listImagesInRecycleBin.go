package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_listImagesInRecycleBinCmd = &cobra.Command{
	Use:   "list-images-in-recycle-bin",
	Short: "Lists one or more AMIs that are currently in the Recycle Bin.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_listImagesInRecycleBinCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_listImagesInRecycleBinCmd).Standalone()

		ec2_listImagesInRecycleBinCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_listImagesInRecycleBinCmd.Flags().String("image-ids", "", "The IDs of the AMIs to list.")
		ec2_listImagesInRecycleBinCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_listImagesInRecycleBinCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_listImagesInRecycleBinCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_listImagesInRecycleBinCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_listImagesInRecycleBinCmd)
}
