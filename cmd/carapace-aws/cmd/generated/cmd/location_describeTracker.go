package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_describeTrackerCmd = &cobra.Command{
	Use:   "describe-tracker",
	Short: "Retrieves the tracker resource details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_describeTrackerCmd).Standalone()

	location_describeTrackerCmd.Flags().String("tracker-name", "", "The name of the tracker resource.")
	location_describeTrackerCmd.MarkFlagRequired("tracker-name")
	locationCmd.AddCommand(location_describeTrackerCmd)
}
