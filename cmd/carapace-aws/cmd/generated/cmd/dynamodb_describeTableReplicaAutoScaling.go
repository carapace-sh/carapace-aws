package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeTableReplicaAutoScalingCmd = &cobra.Command{
	Use:   "describe-table-replica-auto-scaling",
	Short: "Describes auto scaling settings across replicas of the global table at once.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeTableReplicaAutoScalingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_describeTableReplicaAutoScalingCmd).Standalone()

		dynamodb_describeTableReplicaAutoScalingCmd.Flags().String("table-name", "", "The name of the table.")
		dynamodb_describeTableReplicaAutoScalingCmd.MarkFlagRequired("table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_describeTableReplicaAutoScalingCmd)
}
