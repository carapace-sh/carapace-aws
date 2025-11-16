package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getRealtimeLogConfig2020_05_31Cmd = &cobra.Command{
	Use:   "get-realtime-log-config2020_05_31",
	Short: "Gets a real-time log configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getRealtimeLogConfig2020_05_31Cmd).Standalone()

	cloudfront_getRealtimeLogConfig2020_05_31Cmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the real-time log configuration to get.")
	cloudfront_getRealtimeLogConfig2020_05_31Cmd.Flags().String("name", "", "The name of the real-time log configuration to get.")
	cloudfrontCmd.AddCommand(cloudfront_getRealtimeLogConfig2020_05_31Cmd)
}
