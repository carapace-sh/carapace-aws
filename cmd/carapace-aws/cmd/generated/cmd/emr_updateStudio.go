package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_updateStudioCmd = &cobra.Command{
	Use:   "update-studio",
	Short: "Updates an Amazon EMR Studio configuration, including attributes such as name, description, and subnets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_updateStudioCmd).Standalone()

	emr_updateStudioCmd.Flags().String("default-s3-location", "", "The Amazon S3 location to back up Workspaces and notebook files for the Amazon EMR Studio.")
	emr_updateStudioCmd.Flags().String("description", "", "A detailed description to assign to the Amazon EMR Studio.")
	emr_updateStudioCmd.Flags().String("encryption-key-arn", "", "The KMS key identifier (ARN) used to encrypt Amazon EMR Studio workspace and notebook files when backed up to Amazon S3.")
	emr_updateStudioCmd.Flags().String("name", "", "A descriptive name for the Amazon EMR Studio.")
	emr_updateStudioCmd.Flags().String("studio-id", "", "The ID of the Amazon EMR Studio to update.")
	emr_updateStudioCmd.Flags().String("subnet-ids", "", "A list of subnet IDs to associate with the Amazon EMR Studio.")
	emr_updateStudioCmd.MarkFlagRequired("studio-id")
	emrCmd.AddCommand(emr_updateStudioCmd)
}
