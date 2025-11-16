package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_createBlueGreenDeploymentCmd = &cobra.Command{
	Use:   "create-blue-green-deployment",
	Short: "Creates a blue/green deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_createBlueGreenDeploymentCmd).Standalone()

	rds_createBlueGreenDeploymentCmd.Flags().String("blue-green-deployment-name", "", "The name of the blue/green deployment.")
	rds_createBlueGreenDeploymentCmd.Flags().String("source", "", "The Amazon Resource Name (ARN) of the source production database.")
	rds_createBlueGreenDeploymentCmd.Flags().String("tags", "", "Tags to assign to the blue/green deployment.")
	rds_createBlueGreenDeploymentCmd.Flags().String("target-allocated-storage", "", "The amount of storage in gibibytes (GiB) to allocate for the green DB instance.")
	rds_createBlueGreenDeploymentCmd.Flags().String("target-dbcluster-parameter-group-name", "", "The DB cluster parameter group associated with the Aurora DB cluster in the green environment.")
	rds_createBlueGreenDeploymentCmd.Flags().String("target-dbinstance-class", "", "Specify the DB instance class for the databases in the green environment.")
	rds_createBlueGreenDeploymentCmd.Flags().String("target-dbparameter-group-name", "", "The DB parameter group associated with the DB instance in the green environment.")
	rds_createBlueGreenDeploymentCmd.Flags().String("target-engine-version", "", "The engine version of the database in the green environment.")
	rds_createBlueGreenDeploymentCmd.Flags().String("target-iops", "", "The amount of Provisioned IOPS (input/output operations per second) to allocate for the green DB instance.")
	rds_createBlueGreenDeploymentCmd.Flags().String("target-storage-throughput", "", "The storage throughput value for the green DB instance.")
	rds_createBlueGreenDeploymentCmd.Flags().String("target-storage-type", "", "The storage type to associate with the green DB instance.")
	rds_createBlueGreenDeploymentCmd.Flags().String("upgrade-target-storage-config", "", "Whether to upgrade the storage file system configuration on the green database.")
	rds_createBlueGreenDeploymentCmd.MarkFlagRequired("blue-green-deployment-name")
	rds_createBlueGreenDeploymentCmd.MarkFlagRequired("source")
	rdsCmd.AddCommand(rds_createBlueGreenDeploymentCmd)
}
