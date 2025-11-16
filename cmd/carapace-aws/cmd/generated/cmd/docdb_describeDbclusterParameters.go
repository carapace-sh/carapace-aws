package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeDbclusterParametersCmd = &cobra.Command{
	Use:   "describe-dbcluster-parameters",
	Short: "Returns the detailed parameter list for a particular cluster parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeDbclusterParametersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_describeDbclusterParametersCmd).Standalone()

		docdb_describeDbclusterParametersCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of a specific cluster parameter group to return parameter details for.")
		docdb_describeDbclusterParametersCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		docdb_describeDbclusterParametersCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		docdb_describeDbclusterParametersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		docdb_describeDbclusterParametersCmd.Flags().String("source", "", "A value that indicates to return only parameters for a specific source.")
		docdb_describeDbclusterParametersCmd.MarkFlagRequired("dbcluster-parameter-group-name")
	})
	docdbCmd.AddCommand(docdb_describeDbclusterParametersCmd)
}
