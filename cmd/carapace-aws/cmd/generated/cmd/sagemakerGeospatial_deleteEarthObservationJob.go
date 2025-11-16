package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_deleteEarthObservationJobCmd = &cobra.Command{
	Use:   "delete-earth-observation-job",
	Short: "Use this operation to delete an Earth Observation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_deleteEarthObservationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_deleteEarthObservationJobCmd).Standalone()

		sagemakerGeospatial_deleteEarthObservationJobCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Earth Observation job being deleted.")
		sagemakerGeospatial_deleteEarthObservationJobCmd.MarkFlagRequired("arn")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_deleteEarthObservationJobCmd)
}
