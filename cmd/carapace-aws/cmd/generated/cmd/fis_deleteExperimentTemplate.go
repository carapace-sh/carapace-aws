package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_deleteExperimentTemplateCmd = &cobra.Command{
	Use:   "delete-experiment-template",
	Short: "Deletes the specified experiment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_deleteExperimentTemplateCmd).Standalone()

	fis_deleteExperimentTemplateCmd.Flags().String("id", "", "The ID of the experiment template.")
	fis_deleteExperimentTemplateCmd.MarkFlagRequired("id")
	fisCmd.AddCommand(fis_deleteExperimentTemplateCmd)
}
