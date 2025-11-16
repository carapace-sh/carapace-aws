package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_deleteControlCmd = &cobra.Command{
	Use:   "delete-control",
	Short: "Deletes a custom control in Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_deleteControlCmd).Standalone()

	auditmanager_deleteControlCmd.Flags().String("control-id", "", "The unique identifier for the control.")
	auditmanager_deleteControlCmd.MarkFlagRequired("control-id")
	auditmanagerCmd.AddCommand(auditmanager_deleteControlCmd)
}
