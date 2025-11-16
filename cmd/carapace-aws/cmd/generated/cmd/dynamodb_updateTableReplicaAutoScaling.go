package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_updateTableReplicaAutoScalingCmd = &cobra.Command{
	Use:   "update-table-replica-auto-scaling",
	Short: "Updates auto scaling settings on your global tables at once.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_updateTableReplicaAutoScalingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_updateTableReplicaAutoScalingCmd).Standalone()

		dynamodb_updateTableReplicaAutoScalingCmd.Flags().String("global-secondary-index-updates", "", "Represents the auto scaling settings of the global secondary indexes of the replica to be updated.")
		dynamodb_updateTableReplicaAutoScalingCmd.Flags().String("provisioned-write-capacity-auto-scaling-update", "", "")
		dynamodb_updateTableReplicaAutoScalingCmd.Flags().String("replica-updates", "", "Represents the auto scaling settings of replicas of the table that will be modified.")
		dynamodb_updateTableReplicaAutoScalingCmd.Flags().String("table-name", "", "The name of the global table to be updated.")
		dynamodb_updateTableReplicaAutoScalingCmd.MarkFlagRequired("table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_updateTableReplicaAutoScalingCmd)
}
