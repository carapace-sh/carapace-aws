package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbsubnetGroupCmd = &cobra.Command{
	Use:   "delete-dbsubnet-group",
	Short: "Deletes a DB subnet group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbsubnetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteDbsubnetGroupCmd).Standalone()

		rds_deleteDbsubnetGroupCmd.Flags().String("dbsubnet-group-name", "", "The name of the database subnet group to delete.")
		rds_deleteDbsubnetGroupCmd.MarkFlagRequired("dbsubnet-group-name")
	})
	rdsCmd.AddCommand(rds_deleteDbsubnetGroupCmd)
}
