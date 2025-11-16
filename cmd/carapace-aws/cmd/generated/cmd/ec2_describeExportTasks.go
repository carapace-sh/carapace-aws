package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeExportTasksCmd = &cobra.Command{
	Use:   "describe-export-tasks",
	Short: "Describes the specified export instance tasks or all of your export instance tasks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeExportTasksCmd).Standalone()

	ec2_describeExportTasksCmd.Flags().String("export-task-ids", "", "The export task IDs.")
	ec2_describeExportTasksCmd.Flags().String("filters", "", "the filters for the export tasks.")
	ec2Cmd.AddCommand(ec2_describeExportTasksCmd)
}
