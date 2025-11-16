package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerFeaturestoreRuntime_getRecordCmd = &cobra.Command{
	Use:   "get-record",
	Short: "Use for `OnlineStore` serving from a `FeatureStore`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerFeaturestoreRuntime_getRecordCmd).Standalone()

	sagemakerFeaturestoreRuntime_getRecordCmd.Flags().String("expiration-time-response", "", "Parameter to request `ExpiresAt` in response.")
	sagemakerFeaturestoreRuntime_getRecordCmd.Flags().String("feature-group-name", "", "The name or Amazon Resource Name (ARN) of the feature group from which you want to retrieve a record.")
	sagemakerFeaturestoreRuntime_getRecordCmd.Flags().String("feature-names", "", "List of names of Features to be retrieved.")
	sagemakerFeaturestoreRuntime_getRecordCmd.Flags().String("record-identifier-value-as-string", "", "The value that corresponds to `RecordIdentifier` type and uniquely identifies the record in the `FeatureGroup`.")
	sagemakerFeaturestoreRuntime_getRecordCmd.MarkFlagRequired("feature-group-name")
	sagemakerFeaturestoreRuntime_getRecordCmd.MarkFlagRequired("record-identifier-value-as-string")
	sagemakerFeaturestoreRuntimeCmd.AddCommand(sagemakerFeaturestoreRuntime_getRecordCmd)
}
