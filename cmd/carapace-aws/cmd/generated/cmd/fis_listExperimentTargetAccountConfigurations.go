package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_listExperimentTargetAccountConfigurationsCmd = &cobra.Command{
	Use:   "list-experiment-target-account-configurations",
	Short: "Lists the target account configurations of the specified experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_listExperimentTargetAccountConfigurationsCmd).Standalone()

	fis_listExperimentTargetAccountConfigurationsCmd.Flags().String("experiment-id", "", "The ID of the experiment.")
	fis_listExperimentTargetAccountConfigurationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	fis_listExperimentTargetAccountConfigurationsCmd.MarkFlagRequired("experiment-id")
	fisCmd.AddCommand(fis_listExperimentTargetAccountConfigurationsCmd)
}
