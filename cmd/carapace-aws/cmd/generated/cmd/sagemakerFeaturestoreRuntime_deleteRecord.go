package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerFeaturestoreRuntime_deleteRecordCmd = &cobra.Command{
	Use:   "delete-record",
	Short: "Deletes a `Record` from a `FeatureGroup` in the `OnlineStore`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerFeaturestoreRuntime_deleteRecordCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerFeaturestoreRuntime_deleteRecordCmd).Standalone()

		sagemakerFeaturestoreRuntime_deleteRecordCmd.Flags().String("deletion-mode", "", "The name of the deletion mode for deleting the record.")
		sagemakerFeaturestoreRuntime_deleteRecordCmd.Flags().String("event-time", "", "Timestamp indicating when the deletion event occurred.")
		sagemakerFeaturestoreRuntime_deleteRecordCmd.Flags().String("feature-group-name", "", "The name or Amazon Resource Name (ARN) of the feature group to delete the record from.")
		sagemakerFeaturestoreRuntime_deleteRecordCmd.Flags().String("record-identifier-value-as-string", "", "The value for the `RecordIdentifier` that uniquely identifies the record, in string format.")
		sagemakerFeaturestoreRuntime_deleteRecordCmd.Flags().String("target-stores", "", "A list of stores from which you're deleting the record.")
		sagemakerFeaturestoreRuntime_deleteRecordCmd.MarkFlagRequired("event-time")
		sagemakerFeaturestoreRuntime_deleteRecordCmd.MarkFlagRequired("feature-group-name")
		sagemakerFeaturestoreRuntime_deleteRecordCmd.MarkFlagRequired("record-identifier-value-as-string")
	})
	sagemakerFeaturestoreRuntimeCmd.AddCommand(sagemakerFeaturestoreRuntime_deleteRecordCmd)
}
