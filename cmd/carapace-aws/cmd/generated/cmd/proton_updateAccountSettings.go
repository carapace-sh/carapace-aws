package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateAccountSettingsCmd = &cobra.Command{
	Use:   "update-account-settings",
	Short: "Update Proton settings that are used for multiple services in the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateAccountSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_updateAccountSettingsCmd).Standalone()

		proton_updateAccountSettingsCmd.Flags().Bool("delete-pipeline-provisioning-repository", false, "Set to `true` to remove a configured pipeline repository from the account settings.")
		proton_updateAccountSettingsCmd.Flags().Bool("no-delete-pipeline-provisioning-repository", false, "Set to `true` to remove a configured pipeline repository from the account settings.")
		proton_updateAccountSettingsCmd.Flags().String("pipeline-codebuild-role-arn", "", "The Amazon Resource Name (ARN) of the service role you want to use for provisioning pipelines.")
		proton_updateAccountSettingsCmd.Flags().String("pipeline-provisioning-repository", "", "A linked repository for pipeline provisioning.")
		proton_updateAccountSettingsCmd.Flags().String("pipeline-service-role-arn", "", "The Amazon Resource Name (ARN) of the service role you want to use for provisioning pipelines.")
		proton_updateAccountSettingsCmd.Flag("no-delete-pipeline-provisioning-repository").Hidden = true
	})
	protonCmd.AddCommand(proton_updateAccountSettingsCmd)
}
