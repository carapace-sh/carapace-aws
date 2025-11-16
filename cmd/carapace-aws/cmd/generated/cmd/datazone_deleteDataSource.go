package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteDataSourceCmd = &cobra.Command{
	Use:   "delete-data-source",
	Short: "Deletes a data source in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteDataSourceCmd).Standalone()

		datazone_deleteDataSourceCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_deleteDataSourceCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which the data source is deleted.")
		datazone_deleteDataSourceCmd.Flags().String("identifier", "", "The identifier of the data source that is deleted.")
		datazone_deleteDataSourceCmd.Flags().Bool("no-retain-permissions-on-revoke-failure", false, "Specifies that the granted permissions are retained in case of a self-subscribe functionality failure for a data source.")
		datazone_deleteDataSourceCmd.Flags().Bool("retain-permissions-on-revoke-failure", false, "Specifies that the granted permissions are retained in case of a self-subscribe functionality failure for a data source.")
		datazone_deleteDataSourceCmd.MarkFlagRequired("domain-identifier")
		datazone_deleteDataSourceCmd.MarkFlagRequired("identifier")
		datazone_deleteDataSourceCmd.Flag("no-retain-permissions-on-revoke-failure").Hidden = true
	})
	datazoneCmd.AddCommand(datazone_deleteDataSourceCmd)
}
