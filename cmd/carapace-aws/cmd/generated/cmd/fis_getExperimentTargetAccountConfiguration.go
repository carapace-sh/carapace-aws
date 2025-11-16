package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_getExperimentTargetAccountConfigurationCmd = &cobra.Command{
	Use:   "get-experiment-target-account-configuration",
	Short: "Gets information about the specified target account configuration of the experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_getExperimentTargetAccountConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fis_getExperimentTargetAccountConfigurationCmd).Standalone()

		fis_getExperimentTargetAccountConfigurationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the target account.")
		fis_getExperimentTargetAccountConfigurationCmd.Flags().String("experiment-id", "", "The ID of the experiment.")
		fis_getExperimentTargetAccountConfigurationCmd.MarkFlagRequired("account-id")
		fis_getExperimentTargetAccountConfigurationCmd.MarkFlagRequired("experiment-id")
	})
	fisCmd.AddCommand(fis_getExperimentTargetAccountConfigurationCmd)
}
