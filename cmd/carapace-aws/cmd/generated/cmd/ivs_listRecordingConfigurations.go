package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_listRecordingConfigurationsCmd = &cobra.Command{
	Use:   "list-recording-configurations",
	Short: "Gets summary information about all recording configurations in your account, in the Amazon Web Services region where the API request is processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_listRecordingConfigurationsCmd).Standalone()

	ivs_listRecordingConfigurationsCmd.Flags().String("max-results", "", "Maximum number of recording configurations to return.")
	ivs_listRecordingConfigurationsCmd.Flags().String("next-token", "", "The first recording configuration to retrieve.")
	ivsCmd.AddCommand(ivs_listRecordingConfigurationsCmd)
}
