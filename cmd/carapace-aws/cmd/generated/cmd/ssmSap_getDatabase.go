package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_getDatabaseCmd = &cobra.Command{
	Use:   "get-database",
	Short: "Gets the SAP HANA database of an application registered with AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_getDatabaseCmd).Standalone()

	ssmSap_getDatabaseCmd.Flags().String("application-id", "", "The ID of the application.")
	ssmSap_getDatabaseCmd.Flags().String("component-id", "", "The ID of the component.")
	ssmSap_getDatabaseCmd.Flags().String("database-arn", "", "The Amazon Resource Name (ARN) of the database.")
	ssmSap_getDatabaseCmd.Flags().String("database-id", "", "The ID of the database.")
	ssmSapCmd.AddCommand(ssmSap_getDatabaseCmd)
}
