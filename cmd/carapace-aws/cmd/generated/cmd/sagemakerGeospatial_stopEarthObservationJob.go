package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_stopEarthObservationJobCmd = &cobra.Command{
	Use:   "stop-earth-observation-job",
	Short: "Use this operation to stop an existing earth observation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_stopEarthObservationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_stopEarthObservationJobCmd).Standalone()

		sagemakerGeospatial_stopEarthObservationJobCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Earth Observation job being stopped.")
		sagemakerGeospatial_stopEarthObservationJobCmd.MarkFlagRequired("arn")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_stopEarthObservationJobCmd)
}
