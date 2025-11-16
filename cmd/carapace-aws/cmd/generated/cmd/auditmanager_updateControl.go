package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_updateControlCmd = &cobra.Command{
	Use:   "update-control",
	Short: "Updates a custom control in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_updateControlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_updateControlCmd).Standalone()

		auditmanager_updateControlCmd.Flags().String("action-plan-instructions", "", "The recommended actions to carry out if the control isn't fulfilled.")
		auditmanager_updateControlCmd.Flags().String("action-plan-title", "", "The title of the action plan for remediating the control.")
		auditmanager_updateControlCmd.Flags().String("control-id", "", "The identifier for the control.")
		auditmanager_updateControlCmd.Flags().String("control-mapping-sources", "", "The data mapping sources for the control.")
		auditmanager_updateControlCmd.Flags().String("description", "", "The optional description of the control.")
		auditmanager_updateControlCmd.Flags().String("name", "", "The name of the updated control.")
		auditmanager_updateControlCmd.Flags().String("testing-information", "", "The steps that you should follow to determine if the control is met.")
		auditmanager_updateControlCmd.MarkFlagRequired("control-id")
		auditmanager_updateControlCmd.MarkFlagRequired("control-mapping-sources")
		auditmanager_updateControlCmd.MarkFlagRequired("name")
	})
	auditmanagerCmd.AddCommand(auditmanager_updateControlCmd)
}
