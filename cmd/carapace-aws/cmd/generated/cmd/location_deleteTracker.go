package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_deleteTrackerCmd = &cobra.Command{
	Use:   "delete-tracker",
	Short: "Deletes a tracker resource from your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_deleteTrackerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_deleteTrackerCmd).Standalone()

		location_deleteTrackerCmd.Flags().String("tracker-name", "", "The name of the tracker resource to be deleted.")
		location_deleteTrackerCmd.MarkFlagRequired("tracker-name")
	})
	locationCmd.AddCommand(location_deleteTrackerCmd)
}
