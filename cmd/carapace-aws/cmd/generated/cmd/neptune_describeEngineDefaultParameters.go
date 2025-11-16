package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeEngineDefaultParametersCmd = &cobra.Command{
	Use:   "describe-engine-default-parameters",
	Short: "Returns the default engine and system parameter information for the specified database engine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeEngineDefaultParametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_describeEngineDefaultParametersCmd).Standalone()

		neptune_describeEngineDefaultParametersCmd.Flags().String("dbparameter-group-family", "", "The name of the DB parameter group family.")
		neptune_describeEngineDefaultParametersCmd.Flags().String("filters", "", "Not currently supported.")
		neptune_describeEngineDefaultParametersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeEngineDefaultParameters` request.")
		neptune_describeEngineDefaultParametersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		neptune_describeEngineDefaultParametersCmd.MarkFlagRequired("dbparameter-group-family")
	})
	neptuneCmd.AddCommand(neptune_describeEngineDefaultParametersCmd)
}
