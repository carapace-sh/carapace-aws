package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerFeaturestoreRuntime_putRecordCmd = &cobra.Command{
	Use:   "put-record",
	Short: "The `PutRecord` API is used to ingest a list of `Records` into your feature group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerFeaturestoreRuntime_putRecordCmd).Standalone()

	sagemakerFeaturestoreRuntime_putRecordCmd.Flags().String("feature-group-name", "", "The name or Amazon Resource Name (ARN) of the feature group that you want to insert the record into.")
	sagemakerFeaturestoreRuntime_putRecordCmd.Flags().String("record", "", "List of FeatureValues to be inserted.")
	sagemakerFeaturestoreRuntime_putRecordCmd.Flags().String("target-stores", "", "A list of stores to which you're adding the record.")
	sagemakerFeaturestoreRuntime_putRecordCmd.Flags().String("ttl-duration", "", "Time to live duration, where the record is hard deleted after the expiration time is reached; `ExpiresAt` = `EventTime` + `TtlDuration`.")
	sagemakerFeaturestoreRuntime_putRecordCmd.MarkFlagRequired("feature-group-name")
	sagemakerFeaturestoreRuntime_putRecordCmd.MarkFlagRequired("record")
	sagemakerFeaturestoreRuntimeCmd.AddCommand(sagemakerFeaturestoreRuntime_putRecordCmd)
}
