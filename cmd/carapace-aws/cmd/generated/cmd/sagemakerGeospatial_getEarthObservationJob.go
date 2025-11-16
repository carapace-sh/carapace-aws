package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_getEarthObservationJobCmd = &cobra.Command{
	Use:   "get-earth-observation-job",
	Short: "Get the details for a previously initiated Earth Observation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_getEarthObservationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_getEarthObservationJobCmd).Standalone()

		sagemakerGeospatial_getEarthObservationJobCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Earth Observation job.")
		sagemakerGeospatial_getEarthObservationJobCmd.MarkFlagRequired("arn")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_getEarthObservationJobCmd)
}
