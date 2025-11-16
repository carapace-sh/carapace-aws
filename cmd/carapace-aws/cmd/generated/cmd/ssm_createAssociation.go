package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_createAssociationCmd = &cobra.Command{
	Use:   "create-association",
	Short: "A State Manager association defines the state that you want to maintain on your managed nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_createAssociationCmd).Standalone()

	ssm_createAssociationCmd.Flags().String("alarm-configuration", "", "")
	ssm_createAssociationCmd.Flags().String("apply-only-at-cron-interval", "", "By default, when you create a new association, the system runs it immediately after it is created and then according to the schedule you specified and when target changes are detected.")
	ssm_createAssociationCmd.Flags().String("association-name", "", "Specify a descriptive name for the association.")
	ssm_createAssociationCmd.Flags().String("automation-target-parameter-name", "", "Choose the parameter that will define how your automation will branch out.")
	ssm_createAssociationCmd.Flags().String("calendar-names", "", "The names of Amazon Resource Names (ARNs) of the Change Calendar type documents you want to gate your associations under.")
	ssm_createAssociationCmd.Flags().String("compliance-severity", "", "The severity level to assign to the association.")
	ssm_createAssociationCmd.Flags().String("document-version", "", "The document version you want to associate with the targets.")
	ssm_createAssociationCmd.Flags().String("duration", "", "The number of hours the association can run before it is canceled.")
	ssm_createAssociationCmd.Flags().String("instance-id", "", "The managed node ID.")
	ssm_createAssociationCmd.Flags().String("max-concurrency", "", "The maximum number of targets allowed to run the association at the same time.")
	ssm_createAssociationCmd.Flags().String("max-errors", "", "The number of errors that are allowed before the system stops sending requests to run the association on additional targets.")
	ssm_createAssociationCmd.Flags().String("name", "", "The name of the SSM Command document or Automation runbook that contains the configuration information for the managed node.")
	ssm_createAssociationCmd.Flags().String("output-location", "", "An Amazon Simple Storage Service (Amazon S3) bucket where you want to store the output details of the request.")
	ssm_createAssociationCmd.Flags().String("parameters", "", "The parameters for the runtime configuration of the document.")
	ssm_createAssociationCmd.Flags().String("schedule-expression", "", "A cron expression when the association will be applied to the targets.")
	ssm_createAssociationCmd.Flags().String("schedule-offset", "", "Number of days to wait after the scheduled day to run an association.")
	ssm_createAssociationCmd.Flags().String("sync-compliance", "", "The mode for generating association compliance.")
	ssm_createAssociationCmd.Flags().String("tags", "", "Adds or overwrites one or more tags for a State Manager association.")
	ssm_createAssociationCmd.Flags().String("target-locations", "", "A location is a combination of Amazon Web Services Regions and Amazon Web Services accounts where you want to run the association.")
	ssm_createAssociationCmd.Flags().String("target-maps", "", "A key-value mapping of document parameters to target resources.")
	ssm_createAssociationCmd.Flags().String("targets", "", "The targets for the association.")
	ssm_createAssociationCmd.MarkFlagRequired("name")
	ssmCmd.AddCommand(ssm_createAssociationCmd)
}
