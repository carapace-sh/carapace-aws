package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_resetServiceSettingCmd = &cobra.Command{
	Use:   "reset-service-setting",
	Short: "`ServiceSetting` is an account-level setting for an Amazon Web Services service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_resetServiceSettingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_resetServiceSettingCmd).Standalone()

		ssm_resetServiceSettingCmd.Flags().String("setting-id", "", "The Amazon Resource Name (ARN) of the service setting to reset.")
		ssm_resetServiceSettingCmd.MarkFlagRequired("setting-id")
	})
	ssmCmd.AddCommand(ssm_resetServiceSettingCmd)
}
