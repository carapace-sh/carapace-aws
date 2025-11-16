package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_createDataSourceFromS3Cmd = &cobra.Command{
	Use:   "create-data-source-from-s3",
	Short: "Creates a `DataSource` object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_createDataSourceFromS3Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_createDataSourceFromS3Cmd).Standalone()

		machinelearning_createDataSourceFromS3Cmd.Flags().String("compute-statistics", "", "The compute statistics for a `DataSource`.")
		machinelearning_createDataSourceFromS3Cmd.Flags().String("data-source-id", "", "A user-supplied identifier that uniquely identifies the `DataSource`.")
		machinelearning_createDataSourceFromS3Cmd.Flags().String("data-source-name", "", "A user-supplied name or description of the `DataSource`.")
		machinelearning_createDataSourceFromS3Cmd.Flags().String("data-spec", "", "The data specification of a `DataSource`:")
		machinelearning_createDataSourceFromS3Cmd.MarkFlagRequired("data-source-id")
		machinelearning_createDataSourceFromS3Cmd.MarkFlagRequired("data-spec")
	})
	machinelearningCmd.AddCommand(machinelearning_createDataSourceFromS3Cmd)
}
