package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_resetClusterParameterGroupCmd = &cobra.Command{
	Use:   "reset-cluster-parameter-group",
	Short: "Sets one or more parameters of the specified parameter group to their default values and sets the source values of the parameters to \"engine-default\".",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_resetClusterParameterGroupCmd).Standalone()

	redshift_resetClusterParameterGroupCmd.Flags().Bool("no-reset-all-parameters", false, "If `true`, all parameters in the specified parameter group will be reset to their default values.")
	redshift_resetClusterParameterGroupCmd.Flags().String("parameter-group-name", "", "The name of the cluster parameter group to be reset.")
	redshift_resetClusterParameterGroupCmd.Flags().String("parameters", "", "An array of names of parameters to be reset.")
	redshift_resetClusterParameterGroupCmd.Flags().Bool("reset-all-parameters", false, "If `true`, all parameters in the specified parameter group will be reset to their default values.")
	redshift_resetClusterParameterGroupCmd.Flag("no-reset-all-parameters").Hidden = true
	redshift_resetClusterParameterGroupCmd.MarkFlagRequired("parameter-group-name")
	redshiftCmd.AddCommand(redshift_resetClusterParameterGroupCmd)
}
