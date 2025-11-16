package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updateServiceSettingCmd = &cobra.Command{
	Use:   "update-service-setting",
	Short: "`ServiceSetting` is an account-level setting for an Amazon Web Services service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updateServiceSettingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_updateServiceSettingCmd).Standalone()

		ssm_updateServiceSettingCmd.Flags().String("setting-id", "", "The Amazon Resource Name (ARN) of the service setting to update.")
		ssm_updateServiceSettingCmd.Flags().String("setting-value", "", "The new value to specify for the service setting.")
		ssm_updateServiceSettingCmd.MarkFlagRequired("setting-id")
		ssm_updateServiceSettingCmd.MarkFlagRequired("setting-value")
	})
	ssmCmd.AddCommand(ssm_updateServiceSettingCmd)
}
