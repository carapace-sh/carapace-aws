package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_createParameterGroupCmd = &cobra.Command{
	Use:   "create-parameter-group",
	Short: "Creates a new parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_createParameterGroupCmd).Standalone()

	dax_createParameterGroupCmd.Flags().String("description", "", "A description of the parameter group.")
	dax_createParameterGroupCmd.Flags().String("parameter-group-name", "", "The name of the parameter group to apply to all of the clusters in this replication group.")
	dax_createParameterGroupCmd.MarkFlagRequired("parameter-group-name")
	daxCmd.AddCommand(dax_createParameterGroupCmd)
}
