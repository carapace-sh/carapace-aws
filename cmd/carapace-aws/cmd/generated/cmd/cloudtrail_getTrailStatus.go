package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_getTrailStatusCmd = &cobra.Command{
	Use:   "get-trail-status",
	Short: "Returns a JSON-formatted list of information about the specified trail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_getTrailStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_getTrailStatusCmd).Standalone()

		cloudtrail_getTrailStatusCmd.Flags().String("name", "", "Specifies the name or the CloudTrail ARN of the trail for which you are requesting status.")
		cloudtrail_getTrailStatusCmd.MarkFlagRequired("name")
	})
	cloudtrailCmd.AddCommand(cloudtrail_getTrailStatusCmd)
}
