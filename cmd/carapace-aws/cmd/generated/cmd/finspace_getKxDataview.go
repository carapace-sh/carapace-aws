package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_getKxDataviewCmd = &cobra.Command{
	Use:   "get-kx-dataview",
	Short: "Retrieves details of the dataview.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_getKxDataviewCmd).Standalone()

	finspace_getKxDataviewCmd.Flags().String("database-name", "", "The name of the database where you created the dataview.")
	finspace_getKxDataviewCmd.Flags().String("dataview-name", "", "A unique identifier for the dataview.")
	finspace_getKxDataviewCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment, from where you want to retrieve the dataview details.")
	finspace_getKxDataviewCmd.MarkFlagRequired("database-name")
	finspace_getKxDataviewCmd.MarkFlagRequired("dataview-name")
	finspace_getKxDataviewCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_getKxDataviewCmd)
}
