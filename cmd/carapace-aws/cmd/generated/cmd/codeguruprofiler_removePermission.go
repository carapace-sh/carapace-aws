package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_removePermissionCmd = &cobra.Command{
	Use:   "remove-permission",
	Short: "Removes permissions from a profiling group's resource-based policy that are provided using an action group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_removePermissionCmd).Standalone()

	codeguruprofiler_removePermissionCmd.Flags().String("action-group", "", "Specifies an action group that contains the permissions to remove from a profiling group's resource-based policy.")
	codeguruprofiler_removePermissionCmd.Flags().String("profiling-group-name", "", "The name of the profiling group.")
	codeguruprofiler_removePermissionCmd.Flags().String("revision-id", "", "A universally unique identifier (UUID) for the revision of the resource-based policy from which you want to remove permissions.")
	codeguruprofiler_removePermissionCmd.MarkFlagRequired("action-group")
	codeguruprofiler_removePermissionCmd.MarkFlagRequired("profiling-group-name")
	codeguruprofiler_removePermissionCmd.MarkFlagRequired("revision-id")
	codeguruprofilerCmd.AddCommand(codeguruprofiler_removePermissionCmd)
}
