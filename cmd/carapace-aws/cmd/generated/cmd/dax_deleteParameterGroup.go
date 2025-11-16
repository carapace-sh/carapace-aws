package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_deleteParameterGroupCmd = &cobra.Command{
	Use:   "delete-parameter-group",
	Short: "Deletes the specified parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_deleteParameterGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dax_deleteParameterGroupCmd).Standalone()

		dax_deleteParameterGroupCmd.Flags().String("parameter-group-name", "", "The name of the parameter group to delete.")
		dax_deleteParameterGroupCmd.MarkFlagRequired("parameter-group-name")
	})
	daxCmd.AddCommand(dax_deleteParameterGroupCmd)
}
