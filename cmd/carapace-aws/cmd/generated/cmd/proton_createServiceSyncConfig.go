package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_createServiceSyncConfigCmd = &cobra.Command{
	Use:   "create-service-sync-config",
	Short: "Create the Proton Ops configuration file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_createServiceSyncConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_createServiceSyncConfigCmd).Standalone()

		proton_createServiceSyncConfigCmd.Flags().String("branch", "", "The repository branch for your Proton Ops file.")
		proton_createServiceSyncConfigCmd.Flags().String("file-path", "", "The path to the Proton Ops file.")
		proton_createServiceSyncConfigCmd.Flags().String("repository-name", "", "The repository name.")
		proton_createServiceSyncConfigCmd.Flags().String("repository-provider", "", "The provider type for your repository.")
		proton_createServiceSyncConfigCmd.Flags().String("service-name", "", "The name of the service the Proton Ops file is for.")
		proton_createServiceSyncConfigCmd.MarkFlagRequired("branch")
		proton_createServiceSyncConfigCmd.MarkFlagRequired("file-path")
		proton_createServiceSyncConfigCmd.MarkFlagRequired("repository-name")
		proton_createServiceSyncConfigCmd.MarkFlagRequired("repository-provider")
		proton_createServiceSyncConfigCmd.MarkFlagRequired("service-name")
	})
	protonCmd.AddCommand(proton_createServiceSyncConfigCmd)
}
