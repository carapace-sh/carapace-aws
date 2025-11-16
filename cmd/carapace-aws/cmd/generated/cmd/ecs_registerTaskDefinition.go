package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_registerTaskDefinitionCmd = &cobra.Command{
	Use:   "register-task-definition",
	Short: "Registers a new task definition from the supplied `family` and `containerDefinitions`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_registerTaskDefinitionCmd).Standalone()

	ecs_registerTaskDefinitionCmd.Flags().String("container-definitions", "", "A list of container definitions in JSON format that describe the different containers that make up your task.")
	ecs_registerTaskDefinitionCmd.Flags().String("cpu", "", "The number of CPU units used by the task.")
	ecs_registerTaskDefinitionCmd.Flags().String("enable-fault-injection", "", "Enables fault injection when you register your task definition and allows for fault injection requests to be accepted from the task's containers.")
	ecs_registerTaskDefinitionCmd.Flags().String("ephemeral-storage", "", "The amount of ephemeral storage to allocate for the task.")
	ecs_registerTaskDefinitionCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the task execution role that grants the Amazon ECS container agent permission to make Amazon Web Services API calls on your behalf.")
	ecs_registerTaskDefinitionCmd.Flags().String("family", "", "You must specify a `family` for a task definition.")
	ecs_registerTaskDefinitionCmd.Flags().String("inference-accelerators", "", "The Elastic Inference accelerators to use for the containers in the task.")
	ecs_registerTaskDefinitionCmd.Flags().String("ipc-mode", "", "The IPC resource namespace to use for the containers in the task.")
	ecs_registerTaskDefinitionCmd.Flags().String("memory", "", "The amount of memory (in MiB) used by the task.")
	ecs_registerTaskDefinitionCmd.Flags().String("network-mode", "", "The Docker networking mode to use for the containers in the task.")
	ecs_registerTaskDefinitionCmd.Flags().String("pid-mode", "", "The process namespace to use for the containers in the task.")
	ecs_registerTaskDefinitionCmd.Flags().String("placement-constraints", "", "An array of placement constraint objects to use for the task.")
	ecs_registerTaskDefinitionCmd.Flags().String("proxy-configuration", "", "The configuration details for the App Mesh proxy.")
	ecs_registerTaskDefinitionCmd.Flags().String("requires-compatibilities", "", "The task launch type that Amazon ECS validates the task definition against.")
	ecs_registerTaskDefinitionCmd.Flags().String("runtime-platform", "", "The operating system that your tasks definitions run on.")
	ecs_registerTaskDefinitionCmd.Flags().String("tags", "", "The metadata that you apply to the task definition to help you categorize and organize them.")
	ecs_registerTaskDefinitionCmd.Flags().String("task-role-arn", "", "The short name or full Amazon Resource Name (ARN) of the IAM role that containers in this task can assume.")
	ecs_registerTaskDefinitionCmd.Flags().String("volumes", "", "A list of volume definitions in JSON format that containers in your task might use.")
	ecs_registerTaskDefinitionCmd.MarkFlagRequired("container-definitions")
	ecs_registerTaskDefinitionCmd.MarkFlagRequired("family")
	ecsCmd.AddCommand(ecs_registerTaskDefinitionCmd)
}
