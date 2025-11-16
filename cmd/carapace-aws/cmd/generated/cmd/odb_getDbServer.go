package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_getDbServerCmd = &cobra.Command{
	Use:   "get-db-server",
	Short: "Returns information about the specified database server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_getDbServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_getDbServerCmd).Standalone()

		odb_getDbServerCmd.Flags().String("cloud-exadata-infrastructure-id", "", "The unique identifier of the Oracle Exadata infrastructure that contains the database server.")
		odb_getDbServerCmd.Flags().String("db-server-id", "", "The unique identifier of the database server to retrieve information about.")
		odb_getDbServerCmd.MarkFlagRequired("cloud-exadata-infrastructure-id")
		odb_getDbServerCmd.MarkFlagRequired("db-server-id")
	})
	odbCmd.AddCommand(odb_getDbServerCmd)
}
