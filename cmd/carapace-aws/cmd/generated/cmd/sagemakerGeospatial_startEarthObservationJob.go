package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_startEarthObservationJobCmd = &cobra.Command{
	Use:   "start-earth-observation-job",
	Short: "Use this operation to create an Earth observation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_startEarthObservationJobCmd).Standalone()

	sagemakerGeospatial_startEarthObservationJobCmd.Flags().String("client-token", "", "A unique token that guarantees that the call to this API is idempotent.")
	sagemakerGeospatial_startEarthObservationJobCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that you specified for the job.")
	sagemakerGeospatial_startEarthObservationJobCmd.Flags().String("input-config", "", "Input configuration information for the Earth Observation job.")
	sagemakerGeospatial_startEarthObservationJobCmd.Flags().String("job-config", "", "An object containing information about the job configuration.")
	sagemakerGeospatial_startEarthObservationJobCmd.Flags().String("kms-key-id", "", "The Key Management Service key ID for server-side encryption.")
	sagemakerGeospatial_startEarthObservationJobCmd.Flags().String("name", "", "The name of the Earth Observation job.")
	sagemakerGeospatial_startEarthObservationJobCmd.Flags().String("tags", "", "Each tag consists of a key and a value.")
	sagemakerGeospatial_startEarthObservationJobCmd.MarkFlagRequired("execution-role-arn")
	sagemakerGeospatial_startEarthObservationJobCmd.MarkFlagRequired("input-config")
	sagemakerGeospatial_startEarthObservationJobCmd.MarkFlagRequired("job-config")
	sagemakerGeospatial_startEarthObservationJobCmd.MarkFlagRequired("name")
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_startEarthObservationJobCmd)
}
