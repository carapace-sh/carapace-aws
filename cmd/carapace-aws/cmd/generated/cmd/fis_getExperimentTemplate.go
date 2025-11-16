package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_getExperimentTemplateCmd = &cobra.Command{
	Use:   "get-experiment-template",
	Short: "Gets information about the specified experiment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_getExperimentTemplateCmd).Standalone()

	fis_getExperimentTemplateCmd.Flags().String("id", "", "The ID of the experiment template.")
	fis_getExperimentTemplateCmd.MarkFlagRequired("id")
	fisCmd.AddCommand(fis_getExperimentTemplateCmd)
}
