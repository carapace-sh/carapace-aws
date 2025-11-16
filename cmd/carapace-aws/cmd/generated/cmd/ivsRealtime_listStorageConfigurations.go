package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_listStorageConfigurationsCmd = &cobra.Command{
	Use:   "list-storage-configurations",
	Short: "Gets summary information about all storage configurations in your account, in the AWS region where the API request is processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_listStorageConfigurationsCmd).Standalone()

	ivsRealtime_listStorageConfigurationsCmd.Flags().String("max-results", "", "Maximum number of storage configurations to return.")
	ivsRealtime_listStorageConfigurationsCmd.Flags().String("next-token", "", "The first storage configuration to retrieve.")
	ivsRealtimeCmd.AddCommand(ivsRealtime_listStorageConfigurationsCmd)
}
