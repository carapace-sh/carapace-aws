package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_updateWorkGroupCmd = &cobra.Command{
	Use:   "update-work-group",
	Short: "Updates the workgroup with the specified name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_updateWorkGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_updateWorkGroupCmd).Standalone()

		athena_updateWorkGroupCmd.Flags().String("configuration-updates", "", "Contains configuration updates for an Athena SQL workgroup.")
		athena_updateWorkGroupCmd.Flags().String("description", "", "The workgroup description.")
		athena_updateWorkGroupCmd.Flags().String("state", "", "The workgroup state that will be updated for the given workgroup.")
		athena_updateWorkGroupCmd.Flags().String("work-group", "", "The specified workgroup that will be updated.")
		athena_updateWorkGroupCmd.MarkFlagRequired("work-group")
	})
	athenaCmd.AddCommand(athena_updateWorkGroupCmd)
}
