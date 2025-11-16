package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_deleteTargetAccountConfigurationCmd = &cobra.Command{
	Use:   "delete-target-account-configuration",
	Short: "Deletes the specified target account configuration of the experiment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_deleteTargetAccountConfigurationCmd).Standalone()

	fis_deleteTargetAccountConfigurationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID of the target account.")
	fis_deleteTargetAccountConfigurationCmd.Flags().String("experiment-template-id", "", "The ID of the experiment template.")
	fis_deleteTargetAccountConfigurationCmd.MarkFlagRequired("account-id")
	fis_deleteTargetAccountConfigurationCmd.MarkFlagRequired("experiment-template-id")
	fisCmd.AddCommand(fis_deleteTargetAccountConfigurationCmd)
}
