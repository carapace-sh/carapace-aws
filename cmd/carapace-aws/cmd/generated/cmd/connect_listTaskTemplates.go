package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listTaskTemplatesCmd = &cobra.Command{
	Use:   "list-task-templates",
	Short: "Lists task templates for the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listTaskTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listTaskTemplatesCmd).Standalone()

		connect_listTaskTemplatesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_listTaskTemplatesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listTaskTemplatesCmd.Flags().String("name", "", "The name of the task template.")
		connect_listTaskTemplatesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_listTaskTemplatesCmd.Flags().String("status", "", "Marks a template as `ACTIVE` or `INACTIVE` for a task to refer to it.")
		connect_listTaskTemplatesCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_listTaskTemplatesCmd)
}
