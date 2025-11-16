package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_deleteFrameworkCmd = &cobra.Command{
	Use:   "delete-framework",
	Short: "Deletes the framework specified by a framework name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_deleteFrameworkCmd).Standalone()

	backup_deleteFrameworkCmd.Flags().String("framework-name", "", "The unique name of a framework.")
	backup_deleteFrameworkCmd.MarkFlagRequired("framework-name")
	backupCmd.AddCommand(backup_deleteFrameworkCmd)
}
