package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_updateSoftwareSetCmd = &cobra.Command{
	Use:   "update-software-set",
	Short: "Updates a software set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_updateSoftwareSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesThinClient_updateSoftwareSetCmd).Standalone()

		workspacesThinClient_updateSoftwareSetCmd.Flags().String("id", "", "The ID of the software set to update.")
		workspacesThinClient_updateSoftwareSetCmd.Flags().String("validation-status", "", "An option to define if the software set has been validated.")
		workspacesThinClient_updateSoftwareSetCmd.MarkFlagRequired("id")
		workspacesThinClient_updateSoftwareSetCmd.MarkFlagRequired("validation-status")
	})
	workspacesThinClientCmd.AddCommand(workspacesThinClient_updateSoftwareSetCmd)
}
