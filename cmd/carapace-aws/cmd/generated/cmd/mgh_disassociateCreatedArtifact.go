package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_disassociateCreatedArtifactCmd = &cobra.Command{
	Use:   "disassociate-created-artifact",
	Short: "Disassociates a created artifact of an AWS resource with a migration task performed by a migration tool that was previously associated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_disassociateCreatedArtifactCmd).Standalone()

	mgh_disassociateCreatedArtifactCmd.Flags().String("created-artifact-name", "", "An ARN of the AWS resource related to the migration (e.g., AMI, EC2 instance, RDS instance, etc.)")
	mgh_disassociateCreatedArtifactCmd.Flags().String("dry-run", "", "Optional boolean flag to indicate whether any effect should take place.")
	mgh_disassociateCreatedArtifactCmd.Flags().String("migration-task-name", "", "Unique identifier that references the migration task to be disassociated with the artifact.")
	mgh_disassociateCreatedArtifactCmd.Flags().String("progress-update-stream", "", "The name of the ProgressUpdateStream.")
	mgh_disassociateCreatedArtifactCmd.MarkFlagRequired("created-artifact-name")
	mgh_disassociateCreatedArtifactCmd.MarkFlagRequired("migration-task-name")
	mgh_disassociateCreatedArtifactCmd.MarkFlagRequired("progress-update-stream")
	mghCmd.AddCommand(mgh_disassociateCreatedArtifactCmd)
}
