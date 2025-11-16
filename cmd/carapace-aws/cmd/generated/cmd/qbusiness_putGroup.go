package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_putGroupCmd = &cobra.Command{
	Use:   "put-group",
	Short: "Create, or updates, a mapping of users—who have access to a document—to groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_putGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_putGroupCmd).Standalone()

		qbusiness_putGroupCmd.Flags().String("application-id", "", "The identifier of the application in which the user and group mapping belongs.")
		qbusiness_putGroupCmd.Flags().String("data-source-id", "", "The identifier of the data source for which you want to map users to their groups.")
		qbusiness_putGroupCmd.Flags().String("group-members", "", "")
		qbusiness_putGroupCmd.Flags().String("group-name", "", "The list that contains your users or sub groups that belong the same group.")
		qbusiness_putGroupCmd.Flags().String("index-id", "", "The identifier of the index in which you want to map users to their groups.")
		qbusiness_putGroupCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that has access to the S3 file that contains your list of users that belong to a group.")
		qbusiness_putGroupCmd.Flags().String("type", "", "The type of the group.")
		qbusiness_putGroupCmd.MarkFlagRequired("application-id")
		qbusiness_putGroupCmd.MarkFlagRequired("group-members")
		qbusiness_putGroupCmd.MarkFlagRequired("group-name")
		qbusiness_putGroupCmd.MarkFlagRequired("index-id")
		qbusiness_putGroupCmd.MarkFlagRequired("type")
	})
	qbusinessCmd.AddCommand(qbusiness_putGroupCmd)
}
