package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_createDataSourceFromRdsCmd = &cobra.Command{
	Use:   "create-data-source-from-rds",
	Short: "Creates a `DataSource` object from an [Amazon Relational Database Service](http://aws.amazon.com/rds/) (Amazon RDS).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_createDataSourceFromRdsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_createDataSourceFromRdsCmd).Standalone()

		machinelearning_createDataSourceFromRdsCmd.Flags().String("compute-statistics", "", "The compute statistics for a `DataSource`.")
		machinelearning_createDataSourceFromRdsCmd.Flags().String("data-source-id", "", "A user-supplied ID that uniquely identifies the `DataSource`.")
		machinelearning_createDataSourceFromRdsCmd.Flags().String("data-source-name", "", "A user-supplied name or description of the `DataSource`.")
		machinelearning_createDataSourceFromRdsCmd.Flags().String("rdsdata", "", "The data specification of an Amazon RDS `DataSource`:")
		machinelearning_createDataSourceFromRdsCmd.Flags().String("role-arn", "", "The role that Amazon ML assumes on behalf of the user to create and activate a data pipeline in the user's account and copy data using the `SelectSqlQuery` query from Amazon RDS to Amazon S3.")
		machinelearning_createDataSourceFromRdsCmd.MarkFlagRequired("data-source-id")
		machinelearning_createDataSourceFromRdsCmd.MarkFlagRequired("rdsdata")
		machinelearning_createDataSourceFromRdsCmd.MarkFlagRequired("role-arn")
	})
	machinelearningCmd.AddCommand(machinelearning_createDataSourceFromRdsCmd)
}
