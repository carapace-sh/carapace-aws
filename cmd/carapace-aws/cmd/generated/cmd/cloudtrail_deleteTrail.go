package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_deleteTrailCmd = &cobra.Command{
	Use:   "delete-trail",
	Short: "Deletes a trail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_deleteTrailCmd).Standalone()

	cloudtrail_deleteTrailCmd.Flags().String("name", "", "Specifies the name or the CloudTrail ARN of the trail to be deleted.")
	cloudtrail_deleteTrailCmd.MarkFlagRequired("name")
	cloudtrailCmd.AddCommand(cloudtrail_deleteTrailCmd)
}
