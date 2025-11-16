package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_listDbServersCmd = &cobra.Command{
	Use:   "list-db-servers",
	Short: "Returns information about the database servers that belong to the specified Exadata infrastructure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_listDbServersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_listDbServersCmd).Standalone()

		odb_listDbServersCmd.Flags().String("cloud-exadata-infrastructure-id", "", "The unique identifier of the Oracle Exadata infrastructure.")
		odb_listDbServersCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		odb_listDbServersCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		odb_listDbServersCmd.MarkFlagRequired("cloud-exadata-infrastructure-id")
	})
	odbCmd.AddCommand(odb_listDbServersCmd)
}
