package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_exportEarthObservationJobCmd = &cobra.Command{
	Use:   "export-earth-observation-job",
	Short: "Use this operation to export results of an Earth Observation job and optionally source images used as input to the EOJ to an Amazon S3 location.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_exportEarthObservationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_exportEarthObservationJobCmd).Standalone()

		sagemakerGeospatial_exportEarthObservationJobCmd.Flags().String("arn", "", "The input Amazon Resource Name (ARN) of the Earth Observation job being exported.")
		sagemakerGeospatial_exportEarthObservationJobCmd.Flags().String("client-token", "", "A unique token that guarantees that the call to this API is idempotent.")
		sagemakerGeospatial_exportEarthObservationJobCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that you specified for the job.")
		sagemakerGeospatial_exportEarthObservationJobCmd.Flags().Bool("export-source-images", false, "The source images provided to the Earth Observation job being exported.")
		sagemakerGeospatial_exportEarthObservationJobCmd.Flags().Bool("no-export-source-images", false, "The source images provided to the Earth Observation job being exported.")
		sagemakerGeospatial_exportEarthObservationJobCmd.Flags().String("output-config", "", "An object containing information about the output file.")
		sagemakerGeospatial_exportEarthObservationJobCmd.MarkFlagRequired("arn")
		sagemakerGeospatial_exportEarthObservationJobCmd.MarkFlagRequired("execution-role-arn")
		sagemakerGeospatial_exportEarthObservationJobCmd.Flag("no-export-source-images").Hidden = true
		sagemakerGeospatial_exportEarthObservationJobCmd.MarkFlagRequired("output-config")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_exportEarthObservationJobCmd)
}
