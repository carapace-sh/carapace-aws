package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeTagsCmd = &cobra.Command{
	Use:   "describe-tags",
	Short: "Describes the specified tags for your EC2 resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeTagsCmd).Standalone()

		ec2_describeTagsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTagsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeTagsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeTagsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeTagsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeTagsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeTagsCmd)
}
