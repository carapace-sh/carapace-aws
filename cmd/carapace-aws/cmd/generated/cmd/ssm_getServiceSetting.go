package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getServiceSettingCmd = &cobra.Command{
	Use:   "get-service-setting",
	Short: "`ServiceSetting` is an account-level setting for an Amazon Web Services service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getServiceSettingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_getServiceSettingCmd).Standalone()

		ssm_getServiceSettingCmd.Flags().String("setting-id", "", "The ID of the service setting to get.")
		ssm_getServiceSettingCmd.MarkFlagRequired("setting-id")
	})
	ssmCmd.AddCommand(ssm_getServiceSettingCmd)
}
