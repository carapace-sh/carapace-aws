package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getControlCmd = &cobra.Command{
	Use:   "get-control",
	Short: "Gets information about a specified control.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getControlCmd).Standalone()

	auditmanager_getControlCmd.Flags().String("control-id", "", "The identifier for the control.")
	auditmanager_getControlCmd.MarkFlagRequired("control-id")
	auditmanagerCmd.AddCommand(auditmanager_getControlCmd)
}
