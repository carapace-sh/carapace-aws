package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_updateCaseCmd = &cobra.Command{
	Use:   "update-case",
	Short: "Updates an existing case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_updateCaseCmd).Standalone()

	securityIr_updateCaseCmd.Flags().String("actual-incident-start-date", "", "Optional element for UpdateCase to provide content for the incident start date field.")
	securityIr_updateCaseCmd.Flags().String("case-id", "", "Required element for UpdateCase to identify the case ID for updates.")
	securityIr_updateCaseCmd.Flags().String("description", "", "Optional element for UpdateCase to provide content for the description field.")
	securityIr_updateCaseCmd.Flags().String("engagement-type", "", "Optional element for UpdateCase to provide content for the engagement type field.")
	securityIr_updateCaseCmd.Flags().String("impacted-accounts-to-add", "", "Optional element for UpdateCase to provide content to add accounts impacted.")
	securityIr_updateCaseCmd.Flags().String("impacted-accounts-to-delete", "", "Optional element for UpdateCase to provide content to add accounts impacted.")
	securityIr_updateCaseCmd.Flags().String("impacted-aws-regions-to-add", "", "Optional element for UpdateCase to provide content to add regions impacted.")
	securityIr_updateCaseCmd.Flags().String("impacted-aws-regions-to-delete", "", "Optional element for UpdateCase to provide content to remove regions impacted.")
	securityIr_updateCaseCmd.Flags().String("impacted-services-to-add", "", "Optional element for UpdateCase to provide content to add services impacted.")
	securityIr_updateCaseCmd.Flags().String("impacted-services-to-delete", "", "Optional element for UpdateCase to provide content to remove services impacted.")
	securityIr_updateCaseCmd.Flags().String("reported-incident-start-date", "", "Optional element for UpdateCase to provide content for the customer reported incident start date field.")
	securityIr_updateCaseCmd.Flags().String("threat-actor-ip-addresses-to-add", "", "Optional element for UpdateCase to provide content to add additional suspicious IP addresses related to a case.")
	securityIr_updateCaseCmd.Flags().String("threat-actor-ip-addresses-to-delete", "", "Optional element for UpdateCase to provide content to remove suspicious IP addresses from a case.")
	securityIr_updateCaseCmd.Flags().String("title", "", "Optional element for UpdateCase to provide content for the title field.")
	securityIr_updateCaseCmd.Flags().String("watchers-to-add", "", "Optional element for UpdateCase to provide content to add additional watchers to a case.")
	securityIr_updateCaseCmd.Flags().String("watchers-to-delete", "", "Optional element for UpdateCase to provide content to remove existing watchers from a case.")
	securityIr_updateCaseCmd.MarkFlagRequired("case-id")
	securityIrCmd.AddCommand(securityIr_updateCaseCmd)
}
