package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_updateTaskTemplateCmd = &cobra.Command{
	Use:   "update-task-template",
	Short: "Updates details about a specific task template in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_updateTaskTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_updateTaskTemplateCmd).Standalone()

		connect_updateTaskTemplateCmd.Flags().String("constraints", "", "Constraints that are applicable to the fields listed.")
		connect_updateTaskTemplateCmd.Flags().String("contact-flow-id", "", "The identifier of the flow that runs by default when a task is created by referencing this template.")
		connect_updateTaskTemplateCmd.Flags().String("defaults", "", "The default values for fields when a task is created by referencing this template.")
		connect_updateTaskTemplateCmd.Flags().String("description", "", "The description of the task template.")
		connect_updateTaskTemplateCmd.Flags().String("fields", "", "Fields that are part of the template.")
		connect_updateTaskTemplateCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_updateTaskTemplateCmd.Flags().String("name", "", "The name of the task template.")
		connect_updateTaskTemplateCmd.Flags().String("self-assign-flow-id", "", "The ContactFlowId for the flow that will be run if this template is used to create a self-assigned task.")
		connect_updateTaskTemplateCmd.Flags().String("status", "", "Marks a template as `ACTIVE` or `INACTIVE` for a task to refer to it.")
		connect_updateTaskTemplateCmd.Flags().String("task-template-id", "", "A unique identifier for the task template.")
		connect_updateTaskTemplateCmd.MarkFlagRequired("instance-id")
		connect_updateTaskTemplateCmd.MarkFlagRequired("task-template-id")
	})
	connectCmd.AddCommand(connect_updateTaskTemplateCmd)
}
