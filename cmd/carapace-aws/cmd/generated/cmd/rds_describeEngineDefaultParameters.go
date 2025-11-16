package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeEngineDefaultParametersCmd = &cobra.Command{
	Use:   "describe-engine-default-parameters",
	Short: "Returns the default engine and system parameter information for the specified database engine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeEngineDefaultParametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeEngineDefaultParametersCmd).Standalone()

		rds_describeEngineDefaultParametersCmd.Flags().String("dbparameter-group-family", "", "The name of the DB parameter group family.")
		rds_describeEngineDefaultParametersCmd.Flags().String("filters", "", "A filter that specifies one or more parameters to describe.")
		rds_describeEngineDefaultParametersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribeEngineDefaultParameters` request.")
		rds_describeEngineDefaultParametersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeEngineDefaultParametersCmd.MarkFlagRequired("dbparameter-group-family")
	})
	rdsCmd.AddCommand(rds_describeEngineDefaultParametersCmd)
}
