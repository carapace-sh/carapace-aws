package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_updateResourceProfileDetectionsCmd = &cobra.Command{
	Use:   "update-resource-profile-detections",
	Short: "Updates the sensitivity scoring settings for an S3 bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_updateResourceProfileDetectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_updateResourceProfileDetectionsCmd).Standalone()

		macie2_updateResourceProfileDetectionsCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the S3 bucket that the request applies to.")
		macie2_updateResourceProfileDetectionsCmd.Flags().String("suppress-data-identifiers", "", "An array of objects, one for each custom data identifier or managed data identifier that detected a type of sensitive data to exclude from the bucket's score.")
		macie2_updateResourceProfileDetectionsCmd.MarkFlagRequired("resource-arn")
	})
	macie2Cmd.AddCommand(macie2_updateResourceProfileDetectionsCmd)
}
