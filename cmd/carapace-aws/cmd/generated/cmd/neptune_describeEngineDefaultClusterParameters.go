package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeEngineDefaultClusterParametersCmd = &cobra.Command{
	Use:   "describe-engine-default-cluster-parameters",
	Short: "Returns the default engine and system parameter information for the cluster database engine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeEngineDefaultClusterParametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_describeEngineDefaultClusterParametersCmd).Standalone()

		neptune_describeEngineDefaultClusterParametersCmd.Flags().String("dbparameter-group-family", "", "The name of the DB cluster parameter group family to return engine parameter information for.")
		neptune_describeEngineDefaultClusterParametersCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		neptune_describeEngineDefaultClusterParametersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeEngineDefaultClusterParameters` request.")
		neptune_describeEngineDefaultClusterParametersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		neptune_describeEngineDefaultClusterParametersCmd.MarkFlagRequired("dbparameter-group-family")
	})
	neptuneCmd.AddCommand(neptune_describeEngineDefaultClusterParametersCmd)
}
