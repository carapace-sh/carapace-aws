package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_listEncoderConfigurationsCmd = &cobra.Command{
	Use:   "list-encoder-configurations",
	Short: "Gets summary information about all EncoderConfigurations in your account, in the AWS region where the API request is processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_listEncoderConfigurationsCmd).Standalone()

	ivsRealtime_listEncoderConfigurationsCmd.Flags().String("max-results", "", "Maximum number of results to return.")
	ivsRealtime_listEncoderConfigurationsCmd.Flags().String("next-token", "", "The first encoder configuration to retrieve.")
	ivsRealtimeCmd.AddCommand(ivsRealtime_listEncoderConfigurationsCmd)
}
