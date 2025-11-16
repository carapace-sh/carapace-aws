package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_deleteDbsubnetGroupCmd = &cobra.Command{
	Use:   "delete-dbsubnet-group",
	Short: "Deletes a DB subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_deleteDbsubnetGroupCmd).Standalone()

	neptune_deleteDbsubnetGroupCmd.Flags().String("dbsubnet-group-name", "", "The name of the database subnet group to delete.")
	neptune_deleteDbsubnetGroupCmd.MarkFlagRequired("dbsubnet-group-name")
	neptuneCmd.AddCommand(neptune_deleteDbsubnetGroupCmd)
}
