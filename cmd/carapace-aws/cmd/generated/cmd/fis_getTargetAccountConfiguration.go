package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_getTargetAccountConfigurationCmd = &cobra.Command{
	Use:   "get-target-account-configuration",
	Short: "Gets information about the specified target account configuration of the experiment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_getTargetAccountConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fis_getTargetAccountConfigurationCmd).Standalone()

		fis_getTargetAccountConfigurationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the target account.")
		fis_getTargetAccountConfigurationCmd.Flags().String("experiment-template-id", "", "The ID of the experiment template.")
		fis_getTargetAccountConfigurationCmd.MarkFlagRequired("account-id")
		fis_getTargetAccountConfigurationCmd.MarkFlagRequired("experiment-template-id")
	})
	fisCmd.AddCommand(fis_getTargetAccountConfigurationCmd)
}
