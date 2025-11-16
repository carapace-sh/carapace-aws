package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_deleteKxDataviewCmd = &cobra.Command{
	Use:   "delete-kx-dataview",
	Short: "Deletes the specified dataview.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_deleteKxDataviewCmd).Standalone()

	finspace_deleteKxDataviewCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_deleteKxDataviewCmd.Flags().String("database-name", "", "The name of the database whose dataview you want to delete.")
	finspace_deleteKxDataviewCmd.Flags().String("dataview-name", "", "The name of the dataview that you want to delete.")
	finspace_deleteKxDataviewCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment, from where you want to delete the dataview.")
	finspace_deleteKxDataviewCmd.MarkFlagRequired("client-token")
	finspace_deleteKxDataviewCmd.MarkFlagRequired("database-name")
	finspace_deleteKxDataviewCmd.MarkFlagRequired("dataview-name")
	finspace_deleteKxDataviewCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_deleteKxDataviewCmd)
}
