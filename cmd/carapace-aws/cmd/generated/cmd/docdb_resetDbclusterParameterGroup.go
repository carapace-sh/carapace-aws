package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_resetDbclusterParameterGroupCmd = &cobra.Command{
	Use:   "reset-dbcluster-parameter-group",
	Short: "Modifies the parameters of a cluster parameter group to the default value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_resetDbclusterParameterGroupCmd).Standalone()

	docdb_resetDbclusterParameterGroupCmd.Flags().String("dbcluster-parameter-group-name", "", "The name of the cluster parameter group to reset.")
	docdb_resetDbclusterParameterGroupCmd.Flags().Bool("no-reset-all-parameters", false, "A value that is set to `true` to reset all parameters in the cluster parameter group to their default values, and `false` otherwise.")
	docdb_resetDbclusterParameterGroupCmd.Flags().String("parameters", "", "A list of parameter names in the cluster parameter group to reset to the default values.")
	docdb_resetDbclusterParameterGroupCmd.Flags().Bool("reset-all-parameters", false, "A value that is set to `true` to reset all parameters in the cluster parameter group to their default values, and `false` otherwise.")
	docdb_resetDbclusterParameterGroupCmd.MarkFlagRequired("dbcluster-parameter-group-name")
	docdb_resetDbclusterParameterGroupCmd.Flag("no-reset-all-parameters").Hidden = true
	docdbCmd.AddCommand(docdb_resetDbclusterParameterGroupCmd)
}
