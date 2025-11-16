package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createTaskTemplateCmd = &cobra.Command{
	Use:   "create-task-template",
	Short: "Creates a new task template in the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createTaskTemplateCmd).Standalone()

	connect_createTaskTemplateCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	connect_createTaskTemplateCmd.Flags().String("constraints", "", "Constraints that are applicable to the fields listed.")
	connect_createTaskTemplateCmd.Flags().String("contact-flow-id", "", "The identifier of the flow that runs by default when a task is created by referencing this template.")
	connect_createTaskTemplateCmd.Flags().String("defaults", "", "The default values for fields when a task is created by referencing this template.")
	connect_createTaskTemplateCmd.Flags().String("description", "", "The description of the task template.")
	connect_createTaskTemplateCmd.Flags().String("fields", "", "Fields that are part of the template.")
	connect_createTaskTemplateCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createTaskTemplateCmd.Flags().String("name", "", "The name of the task template.")
	connect_createTaskTemplateCmd.Flags().String("self-assign-flow-id", "", "The ContactFlowId for the flow that will be run if this template is used to create a self-assigned task.")
	connect_createTaskTemplateCmd.Flags().String("status", "", "Marks a template as `ACTIVE` or `INACTIVE` for a task to refer to it.")
	connect_createTaskTemplateCmd.MarkFlagRequired("fields")
	connect_createTaskTemplateCmd.MarkFlagRequired("instance-id")
	connect_createTaskTemplateCmd.MarkFlagRequired("name")
	connectCmd.AddCommand(connect_createTaskTemplateCmd)
}
