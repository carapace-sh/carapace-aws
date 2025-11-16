package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createInstanceExportTaskCmd = &cobra.Command{
	Use:   "create-instance-export-task",
	Short: "Exports a running or stopped instance to an Amazon S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createInstanceExportTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_createInstanceExportTaskCmd).Standalone()

		ec2_createInstanceExportTaskCmd.Flags().String("description", "", "A description for the conversion task or the resource being exported.")
		ec2_createInstanceExportTaskCmd.Flags().String("export-to-s3-task", "", "The format and location for an export instance task.")
		ec2_createInstanceExportTaskCmd.Flags().String("instance-id", "", "The ID of the instance.")
		ec2_createInstanceExportTaskCmd.Flags().String("tag-specifications", "", "The tags to apply to the export instance task during creation.")
		ec2_createInstanceExportTaskCmd.Flags().String("target-environment", "", "The target virtualization environment.")
		ec2_createInstanceExportTaskCmd.MarkFlagRequired("export-to-s3-task")
		ec2_createInstanceExportTaskCmd.MarkFlagRequired("instance-id")
		ec2_createInstanceExportTaskCmd.MarkFlagRequired("target-environment")
	})
	ec2Cmd.AddCommand(ec2_createInstanceExportTaskCmd)
}
