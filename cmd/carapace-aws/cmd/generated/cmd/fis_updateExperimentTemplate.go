package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_updateExperimentTemplateCmd = &cobra.Command{
	Use:   "update-experiment-template",
	Short: "Updates the specified experiment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_updateExperimentTemplateCmd).Standalone()

	fis_updateExperimentTemplateCmd.Flags().String("actions", "", "The actions for the experiment.")
	fis_updateExperimentTemplateCmd.Flags().String("description", "", "A description for the template.")
	fis_updateExperimentTemplateCmd.Flags().String("experiment-options", "", "The experiment options for the experiment template.")
	fis_updateExperimentTemplateCmd.Flags().String("experiment-report-configuration", "", "The experiment report configuration for the experiment template.")
	fis_updateExperimentTemplateCmd.Flags().String("id", "", "The ID of the experiment template.")
	fis_updateExperimentTemplateCmd.Flags().String("log-configuration", "", "The configuration for experiment logging.")
	fis_updateExperimentTemplateCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that grants the FIS service permission to perform service actions on your behalf.")
	fis_updateExperimentTemplateCmd.Flags().String("stop-conditions", "", "The stop conditions for the experiment.")
	fis_updateExperimentTemplateCmd.Flags().String("targets", "", "The targets for the experiment.")
	fis_updateExperimentTemplateCmd.MarkFlagRequired("id")
	fisCmd.AddCommand(fis_updateExperimentTemplateCmd)
}
