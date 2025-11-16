package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listRealtimeLogConfigs2020_05_31Cmd = &cobra.Command{
	Use:   "list-realtime-log-configs2020_05_31",
	Short: "Gets a list of real-time log configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listRealtimeLogConfigs2020_05_31Cmd).Standalone()

	cloudfront_listRealtimeLogConfigs2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of real-time log configurations.")
	cloudfront_listRealtimeLogConfigs2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of real-time log configurations that you want in the response.")
	cloudfrontCmd.AddCommand(cloudfront_listRealtimeLogConfigs2020_05_31Cmd)
}
