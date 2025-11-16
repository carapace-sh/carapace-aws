package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_createParameterGroupCmd = &cobra.Command{
	Use:   "create-parameter-group",
	Short: "Creates a new MemoryDB parameter group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_createParameterGroupCmd).Standalone()

	memorydb_createParameterGroupCmd.Flags().String("description", "", "An optional description of the parameter group.")
	memorydb_createParameterGroupCmd.Flags().String("family", "", "The name of the parameter group family that the parameter group can be used with.")
	memorydb_createParameterGroupCmd.Flags().String("parameter-group-name", "", "The name of the parameter group.")
	memorydb_createParameterGroupCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	memorydb_createParameterGroupCmd.MarkFlagRequired("family")
	memorydb_createParameterGroupCmd.MarkFlagRequired("parameter-group-name")
	memorydbCmd.AddCommand(memorydb_createParameterGroupCmd)
}
