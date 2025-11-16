package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmSap_updateApplicationSettingsCmd = &cobra.Command{
	Use:   "update-application-settings",
	Short: "Updates the settings of an application registered with AWS Systems Manager for SAP.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmSap_updateApplicationSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmSap_updateApplicationSettingsCmd).Standalone()

		ssmSap_updateApplicationSettingsCmd.Flags().String("application-id", "", "The ID of the application.")
		ssmSap_updateApplicationSettingsCmd.Flags().String("backint", "", "Installation of AWS Backint Agent for SAP HANA.")
		ssmSap_updateApplicationSettingsCmd.Flags().String("credentials-to-add-or-update", "", "The credentials to be added or updated.")
		ssmSap_updateApplicationSettingsCmd.Flags().String("credentials-to-remove", "", "The credentials to be removed.")
		ssmSap_updateApplicationSettingsCmd.Flags().String("database-arn", "", "The Amazon Resource Name of the SAP HANA database that replaces the current SAP HANA connection with the SAP\\_ABAP application.")
		ssmSap_updateApplicationSettingsCmd.MarkFlagRequired("application-id")
	})
	ssmSapCmd.AddCommand(ssmSap_updateApplicationSettingsCmd)
}
