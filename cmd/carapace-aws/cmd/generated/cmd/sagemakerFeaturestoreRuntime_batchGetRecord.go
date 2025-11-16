package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerFeaturestoreRuntime_batchGetRecordCmd = &cobra.Command{
	Use:   "batch-get-record",
	Short: "Retrieves a batch of `Records` from a `FeatureGroup`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerFeaturestoreRuntime_batchGetRecordCmd).Standalone()

	sagemakerFeaturestoreRuntime_batchGetRecordCmd.Flags().String("expiration-time-response", "", "Parameter to request `ExpiresAt` in response.")
	sagemakerFeaturestoreRuntime_batchGetRecordCmd.Flags().String("identifiers", "", "A list containing the name or Amazon Resource Name (ARN) of the `FeatureGroup`, the list of names of `Feature`s to be retrieved, and the corresponding `RecordIdentifier` values as strings.")
	sagemakerFeaturestoreRuntime_batchGetRecordCmd.MarkFlagRequired("identifiers")
	sagemakerFeaturestoreRuntimeCmd.AddCommand(sagemakerFeaturestoreRuntime_batchGetRecordCmd)
}
