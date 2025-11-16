package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_updateServiceSyncConfigCmd = &cobra.Command{
	Use:   "update-service-sync-config",
	Short: "Update the Proton Ops config file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_updateServiceSyncConfigCmd).Standalone()

	proton_updateServiceSyncConfigCmd.Flags().String("branch", "", "The name of the code repository branch where the Proton Ops file is found.")
	proton_updateServiceSyncConfigCmd.Flags().String("file-path", "", "The path to the Proton Ops file.")
	proton_updateServiceSyncConfigCmd.Flags().String("repository-name", "", "The name of the repository where the Proton Ops file is found.")
	proton_updateServiceSyncConfigCmd.Flags().String("repository-provider", "", "The name of the repository provider where the Proton Ops file is found.")
	proton_updateServiceSyncConfigCmd.Flags().String("service-name", "", "The name of the service the Proton Ops file is for.")
	proton_updateServiceSyncConfigCmd.MarkFlagRequired("branch")
	proton_updateServiceSyncConfigCmd.MarkFlagRequired("file-path")
	proton_updateServiceSyncConfigCmd.MarkFlagRequired("repository-name")
	proton_updateServiceSyncConfigCmd.MarkFlagRequired("repository-provider")
	proton_updateServiceSyncConfigCmd.MarkFlagRequired("service-name")
	protonCmd.AddCommand(proton_updateServiceSyncConfigCmd)
}
