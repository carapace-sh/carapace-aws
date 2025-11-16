package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_updateTrackerCmd = &cobra.Command{
	Use:   "update-tracker",
	Short: "Updates the specified properties of a given tracker resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_updateTrackerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_updateTrackerCmd).Standalone()

		location_updateTrackerCmd.Flags().String("description", "", "Updates the description for the tracker resource.")
		location_updateTrackerCmd.Flags().Bool("event-bridge-enabled", false, "Whether to enable position `UPDATE` events from this tracker to be sent to EventBridge.")
		location_updateTrackerCmd.Flags().Bool("kms-key-enable-geospatial-queries", false, "Enables `GeospatialQueries` for a tracker that uses a [Amazon Web Services KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html).")
		location_updateTrackerCmd.Flags().Bool("no-event-bridge-enabled", false, "Whether to enable position `UPDATE` events from this tracker to be sent to EventBridge.")
		location_updateTrackerCmd.Flags().Bool("no-kms-key-enable-geospatial-queries", false, "Enables `GeospatialQueries` for a tracker that uses a [Amazon Web Services KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html).")
		location_updateTrackerCmd.Flags().String("position-filtering", "", "Updates the position filtering for the tracker resource.")
		location_updateTrackerCmd.Flags().String("pricing-plan", "", "No longer used.")
		location_updateTrackerCmd.Flags().String("pricing-plan-data-source", "", "This parameter is no longer used.")
		location_updateTrackerCmd.Flags().String("tracker-name", "", "The name of the tracker resource to update.")
		location_updateTrackerCmd.Flag("no-event-bridge-enabled").Hidden = true
		location_updateTrackerCmd.Flag("no-kms-key-enable-geospatial-queries").Hidden = true
		location_updateTrackerCmd.MarkFlagRequired("tracker-name")
	})
	locationCmd.AddCommand(location_updateTrackerCmd)
}
