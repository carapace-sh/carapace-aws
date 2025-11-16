package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_createSegmentSnapshotCmd = &cobra.Command{
	Use:   "create-segment-snapshot",
	Short: "Triggers a job to export a segment to a specified destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_createSegmentSnapshotCmd).Standalone()

	customerProfiles_createSegmentSnapshotCmd.Flags().String("data-format", "", "The format in which the segment will be exported.")
	customerProfiles_createSegmentSnapshotCmd.Flags().String("destination-uri", "", "The destination to which the segment will be exported.")
	customerProfiles_createSegmentSnapshotCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_createSegmentSnapshotCmd.Flags().String("encryption-key", "", "The Amazon Resource Name (ARN) of the KMS key used to encrypt the exported segment.")
	customerProfiles_createSegmentSnapshotCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that allows Customer Profiles service principal to assume the role for conducting KMS and S3 operations.")
	customerProfiles_createSegmentSnapshotCmd.Flags().String("segment-definition-name", "", "The name of the segment definition used in this snapshot request.")
	customerProfiles_createSegmentSnapshotCmd.MarkFlagRequired("data-format")
	customerProfiles_createSegmentSnapshotCmd.MarkFlagRequired("domain-name")
	customerProfiles_createSegmentSnapshotCmd.MarkFlagRequired("segment-definition-name")
	customerProfilesCmd.AddCommand(customerProfiles_createSegmentSnapshotCmd)
}
