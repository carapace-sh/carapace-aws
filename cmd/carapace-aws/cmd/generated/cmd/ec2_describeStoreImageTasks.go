package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeStoreImageTasksCmd = &cobra.Command{
	Use:   "describe-store-image-tasks",
	Short: "Describes the progress of the AMI store tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeStoreImageTasksCmd).Standalone()

	ec2_describeStoreImageTasksCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeStoreImageTasksCmd.Flags().String("filters", "", "The filters.")
	ec2_describeStoreImageTasksCmd.Flags().String("image-ids", "", "The AMI IDs for which to show progress.")
	ec2_describeStoreImageTasksCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeStoreImageTasksCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	ec2_describeStoreImageTasksCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeStoreImageTasksCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeStoreImageTasksCmd)
}
