package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getTaskTemplateCmd = &cobra.Command{
	Use:   "get-task-template",
	Short: "Gets details about a specific task template in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getTaskTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_getTaskTemplateCmd).Standalone()

		connect_getTaskTemplateCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_getTaskTemplateCmd.Flags().String("snapshot-version", "", "The system generated version of a task template that is associated with a task, when the task is created.")
		connect_getTaskTemplateCmd.Flags().String("task-template-id", "", "A unique identifier for the task template.")
		connect_getTaskTemplateCmd.MarkFlagRequired("instance-id")
		connect_getTaskTemplateCmd.MarkFlagRequired("task-template-id")
	})
	connectCmd.AddCommand(connect_getTaskTemplateCmd)
}
