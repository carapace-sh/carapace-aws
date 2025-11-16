package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruprofiler_putPermissionCmd = &cobra.Command{
	Use:   "put-permission",
	Short: "Adds permissions to a profiling group's resource-based policy that are provided using an action group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruprofiler_putPermissionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruprofiler_putPermissionCmd).Standalone()

		codeguruprofiler_putPermissionCmd.Flags().String("action-group", "", "Specifies an action group that contains permissions to add to a profiling group resource.")
		codeguruprofiler_putPermissionCmd.Flags().String("principals", "", "A list ARNs for the roles and users you want to grant access to the profiling group.")
		codeguruprofiler_putPermissionCmd.Flags().String("profiling-group-name", "", "The name of the profiling group to grant access to.")
		codeguruprofiler_putPermissionCmd.Flags().String("revision-id", "", "A universally unique identifier (UUID) for the revision of the policy you are adding to the profiling group.")
		codeguruprofiler_putPermissionCmd.MarkFlagRequired("action-group")
		codeguruprofiler_putPermissionCmd.MarkFlagRequired("principals")
		codeguruprofiler_putPermissionCmd.MarkFlagRequired("profiling-group-name")
	})
	codeguruprofilerCmd.AddCommand(codeguruprofiler_putPermissionCmd)
}
