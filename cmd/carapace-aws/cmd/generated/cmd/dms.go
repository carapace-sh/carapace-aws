package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dmsCmd = &cobra.Command{
	Use:   "dms",
	Short: "Database Migration Service\n\nDatabase Migration Service (DMS) can migrate your data to and from the most widely used commercial and open-source databases such as Oracle, PostgreSQL, Microsoft SQL Server, Amazon Redshift, MariaDB, Amazon Aurora, MySQL, and SAP Adaptive Server Enterprise (ASE).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dmsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dmsCmd).Standalone()

	})
	rootCmd.AddCommand(dmsCmd)
}
