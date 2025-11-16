package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyClusterParameterGroupCmd = &cobra.Command{
	Use:   "modify-cluster-parameter-group",
	Short: "Modifies the parameters of a parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyClusterParameterGroupCmd).Standalone()

	redshift_modifyClusterParameterGroupCmd.Flags().String("parameter-group-name", "", "The name of the parameter group to be modified.")
	redshift_modifyClusterParameterGroupCmd.Flags().String("parameters", "", "An array of parameters to be modified.")
	redshift_modifyClusterParameterGroupCmd.MarkFlagRequired("parameter-group-name")
	redshift_modifyClusterParameterGroupCmd.MarkFlagRequired("parameters")
	redshiftCmd.AddCommand(redshift_modifyClusterParameterGroupCmd)
}
