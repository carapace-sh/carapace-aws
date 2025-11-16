package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_getTrailCmd = &cobra.Command{
	Use:   "get-trail",
	Short: "Returns settings information for a specified trail.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_getTrailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_getTrailCmd).Standalone()

		cloudtrail_getTrailCmd.Flags().String("name", "", "The name or the Amazon Resource Name (ARN) of the trail for which you want to retrieve settings information.")
		cloudtrail_getTrailCmd.MarkFlagRequired("name")
	})
	cloudtrailCmd.AddCommand(cloudtrail_getTrailCmd)
}
