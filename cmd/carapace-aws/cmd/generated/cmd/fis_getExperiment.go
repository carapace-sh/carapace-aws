package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_getExperimentCmd = &cobra.Command{
	Use:   "get-experiment",
	Short: "Gets information about the specified experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_getExperimentCmd).Standalone()

	fis_getExperimentCmd.Flags().String("id", "", "The ID of the experiment.")
	fis_getExperimentCmd.MarkFlagRequired("id")
	fisCmd.AddCommand(fis_getExperimentCmd)
}
