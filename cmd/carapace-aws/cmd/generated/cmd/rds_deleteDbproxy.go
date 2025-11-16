package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_deleteDbproxyCmd = &cobra.Command{
	Use:   "delete-dbproxy",
	Short: "Deletes an existing DB proxy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_deleteDbproxyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_deleteDbproxyCmd).Standalone()

		rds_deleteDbproxyCmd.Flags().String("dbproxy-name", "", "The name of the DB proxy to delete.")
		rds_deleteDbproxyCmd.MarkFlagRequired("dbproxy-name")
	})
	rdsCmd.AddCommand(rds_deleteDbproxyCmd)
}
