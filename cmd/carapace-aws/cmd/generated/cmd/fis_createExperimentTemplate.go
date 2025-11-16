package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_createExperimentTemplateCmd = &cobra.Command{
	Use:   "create-experiment-template",
	Short: "Creates an experiment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_createExperimentTemplateCmd).Standalone()

	fis_createExperimentTemplateCmd.Flags().String("actions", "", "The actions for the experiment.")
	fis_createExperimentTemplateCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	fis_createExperimentTemplateCmd.Flags().String("description", "", "A description for the experiment template.")
	fis_createExperimentTemplateCmd.Flags().String("experiment-options", "", "The experiment options for the experiment template.")
	fis_createExperimentTemplateCmd.Flags().String("experiment-report-configuration", "", "The experiment report configuration for the experiment template.")
	fis_createExperimentTemplateCmd.Flags().String("log-configuration", "", "The configuration for experiment logging.")
	fis_createExperimentTemplateCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that grants the FIS service permission to perform service actions on your behalf.")
	fis_createExperimentTemplateCmd.Flags().String("stop-conditions", "", "The stop conditions.")
	fis_createExperimentTemplateCmd.Flags().String("tags", "", "The tags to apply to the experiment template.")
	fis_createExperimentTemplateCmd.Flags().String("targets", "", "The targets for the experiment.")
	fis_createExperimentTemplateCmd.MarkFlagRequired("actions")
	fis_createExperimentTemplateCmd.MarkFlagRequired("client-token")
	fis_createExperimentTemplateCmd.MarkFlagRequired("description")
	fis_createExperimentTemplateCmd.MarkFlagRequired("role-arn")
	fis_createExperimentTemplateCmd.MarkFlagRequired("stop-conditions")
	fisCmd.AddCommand(fis_createExperimentTemplateCmd)
}
