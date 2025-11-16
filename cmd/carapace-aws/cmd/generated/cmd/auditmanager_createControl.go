package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_createControlCmd = &cobra.Command{
	Use:   "create-control",
	Short: "Creates a new custom control in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_createControlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_createControlCmd).Standalone()

		auditmanager_createControlCmd.Flags().String("action-plan-instructions", "", "The recommended actions to carry out if the control isn't fulfilled.")
		auditmanager_createControlCmd.Flags().String("action-plan-title", "", "The title of the action plan for remediating the control.")
		auditmanager_createControlCmd.Flags().String("control-mapping-sources", "", "The data mapping sources for the control.")
		auditmanager_createControlCmd.Flags().String("description", "", "The description of the control.")
		auditmanager_createControlCmd.Flags().String("name", "", "The name of the control.")
		auditmanager_createControlCmd.Flags().String("tags", "", "The tags that are associated with the control.")
		auditmanager_createControlCmd.Flags().String("testing-information", "", "The steps to follow to determine if the control is satisfied.")
		auditmanager_createControlCmd.MarkFlagRequired("control-mapping-sources")
		auditmanager_createControlCmd.MarkFlagRequired("name")
	})
	auditmanagerCmd.AddCommand(auditmanager_createControlCmd)
}
