package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeEngineDefaultClusterParametersCmd = &cobra.Command{
	Use:   "describe-engine-default-cluster-parameters",
	Short: "Returns the default engine and system parameter information for the cluster database engine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeEngineDefaultClusterParametersCmd).Standalone()

	rds_describeEngineDefaultClusterParametersCmd.Flags().String("dbparameter-group-family", "", "The name of the DB cluster parameter group family to return engine parameter information for.")
	rds_describeEngineDefaultClusterParametersCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
	rds_describeEngineDefaultClusterParametersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeEngineDefaultClusterParameters` request.")
	rds_describeEngineDefaultClusterParametersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	rds_describeEngineDefaultClusterParametersCmd.MarkFlagRequired("dbparameter-group-family")
	rdsCmd.AddCommand(rds_describeEngineDefaultClusterParametersCmd)
}
