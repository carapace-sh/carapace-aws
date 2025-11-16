package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_listTrailsCmd = &cobra.Command{
	Use:   "list-trails",
	Short: "Lists trails that are in the current account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_listTrailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_listTrailsCmd).Standalone()

		cloudtrail_listTrailsCmd.Flags().String("next-token", "", "The token to use to get the next page of results after a previous API call.")
	})
	cloudtrailCmd.AddCommand(cloudtrail_listTrailsCmd)
}
