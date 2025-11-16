package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_createDataSourceFromRedshiftCmd = &cobra.Command{
	Use:   "create-data-source-from-redshift",
	Short: "Creates a `DataSource` from a database hosted on an Amazon Redshift cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_createDataSourceFromRedshiftCmd).Standalone()

	machinelearning_createDataSourceFromRedshiftCmd.Flags().String("compute-statistics", "", "The compute statistics for a `DataSource`.")
	machinelearning_createDataSourceFromRedshiftCmd.Flags().String("data-source-id", "", "A user-supplied ID that uniquely identifies the `DataSource`.")
	machinelearning_createDataSourceFromRedshiftCmd.Flags().String("data-source-name", "", "A user-supplied name or description of the `DataSource`.")
	machinelearning_createDataSourceFromRedshiftCmd.Flags().String("data-spec", "", "The data specification of an Amazon Redshift `DataSource`:")
	machinelearning_createDataSourceFromRedshiftCmd.Flags().String("role-arn", "", "A fully specified role Amazon Resource Name (ARN).")
	machinelearning_createDataSourceFromRedshiftCmd.MarkFlagRequired("data-source-id")
	machinelearning_createDataSourceFromRedshiftCmd.MarkFlagRequired("data-spec")
	machinelearning_createDataSourceFromRedshiftCmd.MarkFlagRequired("role-arn")
	machinelearningCmd.AddCommand(machinelearning_createDataSourceFromRedshiftCmd)
}
