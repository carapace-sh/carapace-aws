package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateAssociationCmd = &cobra.Command{
	Use:   "update-association",
	Short: "Updates an association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateAssociationCmd).Standalone()

	ssm_updateAssociationCmd.Flags().String("alarm-configuration", "", "")
	ssm_updateAssociationCmd.Flags().String("apply-only-at-cron-interval", "", "By default, when you update an association, the system runs it immediately after it is updated and then according to the schedule you specified.")
	ssm_updateAssociationCmd.Flags().String("association-id", "", "The ID of the association you want to update.")
	ssm_updateAssociationCmd.Flags().String("association-name", "", "The name of the association that you want to update.")
	ssm_updateAssociationCmd.Flags().String("association-version", "", "This parameter is provided for concurrency control purposes.")
	ssm_updateAssociationCmd.Flags().String("automation-target-parameter-name", "", "Choose the parameter that will define how your automation will branch out.")
	ssm_updateAssociationCmd.Flags().String("calendar-names", "", "The names or Amazon Resource Names (ARNs) of the Change Calendar type documents you want to gate your associations under.")
	ssm_updateAssociationCmd.Flags().String("compliance-severity", "", "The severity level to assign to the association.")
	ssm_updateAssociationCmd.Flags().String("document-version", "", "The document version you want update for the association.")
	ssm_updateAssociationCmd.Flags().String("duration", "", "The number of hours the association can run before it is canceled.")
	ssm_updateAssociationCmd.Flags().String("max-concurrency", "", "The maximum number of targets allowed to run the association at the same time.")
	ssm_updateAssociationCmd.Flags().String("max-errors", "", "The number of errors that are allowed before the system stops sending requests to run the association on additional targets.")
	ssm_updateAssociationCmd.Flags().String("name", "", "The name of the SSM Command document or Automation runbook that contains the configuration information for the managed node.")
	ssm_updateAssociationCmd.Flags().String("output-location", "", "An S3 bucket where you want to store the results of this request.")
	ssm_updateAssociationCmd.Flags().String("parameters", "", "The parameters you want to update for the association.")
	ssm_updateAssociationCmd.Flags().String("schedule-expression", "", "The cron expression used to schedule the association that you want to update.")
	ssm_updateAssociationCmd.Flags().String("schedule-offset", "", "Number of days to wait after the scheduled day to run an association.")
	ssm_updateAssociationCmd.Flags().String("sync-compliance", "", "The mode for generating association compliance.")
	ssm_updateAssociationCmd.Flags().String("target-locations", "", "A location is a combination of Amazon Web Services Regions and Amazon Web Services accounts where you want to run the association.")
	ssm_updateAssociationCmd.Flags().String("target-maps", "", "A key-value mapping of document parameters to target resources.")
	ssm_updateAssociationCmd.Flags().String("targets", "", "The targets of the association.")
	ssm_updateAssociationCmd.MarkFlagRequired("association-id")
	ssmCmd.AddCommand(ssm_updateAssociationCmd)
}
