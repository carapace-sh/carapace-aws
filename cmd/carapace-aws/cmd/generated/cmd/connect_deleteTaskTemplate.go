package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteTaskTemplateCmd = &cobra.Command{
	Use:   "delete-task-template",
	Short: "Deletes the task template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteTaskTemplateCmd).Standalone()

	connect_deleteTaskTemplateCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deleteTaskTemplateCmd.Flags().String("task-template-id", "", "A unique identifier for the task template.")
	connect_deleteTaskTemplateCmd.MarkFlagRequired("instance-id")
	connect_deleteTaskTemplateCmd.MarkFlagRequired("task-template-id")
	connectCmd.AddCommand(connect_deleteTaskTemplateCmd)
}
