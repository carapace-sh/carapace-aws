package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_putWorkflowRunPropertiesCmd = &cobra.Command{
	Use:   "put-workflow-run-properties",
	Short: "Puts the specified workflow run properties for the given workflow run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_putWorkflowRunPropertiesCmd).Standalone()

	glue_putWorkflowRunPropertiesCmd.Flags().String("name", "", "Name of the workflow which was run.")
	glue_putWorkflowRunPropertiesCmd.Flags().String("run-id", "", "The ID of the workflow run for which the run properties should be updated.")
	glue_putWorkflowRunPropertiesCmd.Flags().String("run-properties", "", "The properties to put for the specified run.")
	glue_putWorkflowRunPropertiesCmd.MarkFlagRequired("name")
	glue_putWorkflowRunPropertiesCmd.MarkFlagRequired("run-id")
	glue_putWorkflowRunPropertiesCmd.MarkFlagRequired("run-properties")
	glueCmd.AddCommand(glue_putWorkflowRunPropertiesCmd)
}
