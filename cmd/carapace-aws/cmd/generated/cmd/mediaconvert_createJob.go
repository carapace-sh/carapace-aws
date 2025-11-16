package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "Create a new transcoding job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_createJobCmd).Standalone()

	mediaconvert_createJobCmd.Flags().String("acceleration-settings", "", "Optional.")
	mediaconvert_createJobCmd.Flags().String("billing-tags-source", "", "Optionally choose a Billing tags source that AWS Billing and Cost Management will use to display tags for individual output costs on any billing report that you set up.")
	mediaconvert_createJobCmd.Flags().String("client-request-token", "", "Prevent duplicate jobs from being created and ensure idempotency for your requests.")
	mediaconvert_createJobCmd.Flags().String("hop-destinations", "", "Optional.")
	mediaconvert_createJobCmd.Flags().String("job-engine-version", "", "Use Job engine versions to run jobs for your production workflow on one version, while you test and validate the latest version.")
	mediaconvert_createJobCmd.Flags().String("job-template", "", "Optional.")
	mediaconvert_createJobCmd.Flags().String("priority", "", "Optional.")
	mediaconvert_createJobCmd.Flags().String("queue", "", "Optional.")
	mediaconvert_createJobCmd.Flags().String("role", "", "Required.")
	mediaconvert_createJobCmd.Flags().String("settings", "", "JobSettings contains all the transcode settings for a job.")
	mediaconvert_createJobCmd.Flags().String("simulate-reserved-queue", "", "Optional.")
	mediaconvert_createJobCmd.Flags().String("status-update-interval", "", "Optional.")
	mediaconvert_createJobCmd.Flags().String("tags", "", "Optional.")
	mediaconvert_createJobCmd.Flags().String("user-metadata", "", "Optional.")
	mediaconvert_createJobCmd.MarkFlagRequired("role")
	mediaconvert_createJobCmd.MarkFlagRequired("settings")
	mediaconvertCmd.AddCommand(mediaconvert_createJobCmd)
}
