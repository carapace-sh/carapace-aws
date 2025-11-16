package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_updateSettingsCmd = &cobra.Command{
	Use:   "update-settings",
	Short: "Updates Audit Manager settings for the current account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_updateSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_updateSettingsCmd).Standalone()

		auditmanager_updateSettingsCmd.Flags().String("default-assessment-reports-destination", "", "The default S3 destination bucket for storing assessment reports.")
		auditmanager_updateSettingsCmd.Flags().String("default-export-destination", "", "The default S3 destination bucket for storing evidence finder exports.")
		auditmanager_updateSettingsCmd.Flags().String("default-process-owners", "", "A list of the default audit owners.")
		auditmanager_updateSettingsCmd.Flags().String("deregistration-policy", "", "The deregistration policy for your Audit Manager data.")
		auditmanager_updateSettingsCmd.Flags().Bool("evidence-finder-enabled", false, "Specifies whether the evidence finder feature is enabled.")
		auditmanager_updateSettingsCmd.Flags().String("kms-key", "", "The KMS key details.")
		auditmanager_updateSettingsCmd.Flags().Bool("no-evidence-finder-enabled", false, "Specifies whether the evidence finder feature is enabled.")
		auditmanager_updateSettingsCmd.Flags().String("sns-topic", "", "The Amazon Simple Notification Service (Amazon SNS) topic that Audit Manager sends notifications to.")
		auditmanager_updateSettingsCmd.Flag("no-evidence-finder-enabled").Hidden = true
	})
	auditmanagerCmd.AddCommand(auditmanager_updateSettingsCmd)
}
