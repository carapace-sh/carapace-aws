package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeEngineDefaultClusterParametersCmd = &cobra.Command{
	Use:   "describe-engine-default-cluster-parameters",
	Short: "Returns the default engine and system parameter information for the cluster database engine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeEngineDefaultClusterParametersCmd).Standalone()

	docdb_describeEngineDefaultClusterParametersCmd.Flags().String("dbparameter-group-family", "", "The name of the cluster parameter group family to return the engine parameter information for.")
	docdb_describeEngineDefaultClusterParametersCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	docdb_describeEngineDefaultClusterParametersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	docdb_describeEngineDefaultClusterParametersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	docdb_describeEngineDefaultClusterParametersCmd.MarkFlagRequired("dbparameter-group-family")
	docdbCmd.AddCommand(docdb_describeEngineDefaultClusterParametersCmd)
}
