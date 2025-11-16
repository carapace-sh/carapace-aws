package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_updateKxDataviewCmd = &cobra.Command{
	Use:   "update-kx-dataview",
	Short: "Updates the specified dataview.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_updateKxDataviewCmd).Standalone()

	finspace_updateKxDataviewCmd.Flags().String("changeset-id", "", "A unique identifier for the changeset.")
	finspace_updateKxDataviewCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_updateKxDataviewCmd.Flags().String("database-name", "", "The name of the database.")
	finspace_updateKxDataviewCmd.Flags().String("dataview-name", "", "The name of the dataview that you want to update.")
	finspace_updateKxDataviewCmd.Flags().String("description", "", "The description for a dataview.")
	finspace_updateKxDataviewCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment, where you want to update the dataview.")
	finspace_updateKxDataviewCmd.Flags().String("segment-configurations", "", "The configuration that contains the database path of the data that you want to place on each selected volume.")
	finspace_updateKxDataviewCmd.MarkFlagRequired("client-token")
	finspace_updateKxDataviewCmd.MarkFlagRequired("database-name")
	finspace_updateKxDataviewCmd.MarkFlagRequired("dataview-name")
	finspace_updateKxDataviewCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_updateKxDataviewCmd)
}
