package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_associateCreatedArtifactCmd = &cobra.Command{
	Use:   "associate-created-artifact",
	Short: "Associates a created artifact of an AWS cloud resource, the target receiving the migration, with the migration task performed by a migration tool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_associateCreatedArtifactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgh_associateCreatedArtifactCmd).Standalone()

		mgh_associateCreatedArtifactCmd.Flags().String("created-artifact", "", "An ARN of the AWS resource related to the migration (e.g., AMI, EC2 instance, RDS instance, etc.)")
		mgh_associateCreatedArtifactCmd.Flags().String("dry-run", "", "Optional boolean flag to indicate whether any effect should take place.")
		mgh_associateCreatedArtifactCmd.Flags().String("migration-task-name", "", "Unique identifier that references the migration task.")
		mgh_associateCreatedArtifactCmd.Flags().String("progress-update-stream", "", "The name of the ProgressUpdateStream.")
		mgh_associateCreatedArtifactCmd.MarkFlagRequired("created-artifact")
		mgh_associateCreatedArtifactCmd.MarkFlagRequired("migration-task-name")
		mgh_associateCreatedArtifactCmd.MarkFlagRequired("progress-update-stream")
	})
	mghCmd.AddCommand(mgh_associateCreatedArtifactCmd)
}
