package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteRealtimeLogConfig2020_05_31Cmd = &cobra.Command{
	Use:   "delete-realtime-log-config2020_05_31",
	Short: "Deletes a real-time log configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteRealtimeLogConfig2020_05_31Cmd).Standalone()

	cloudfront_deleteRealtimeLogConfig2020_05_31Cmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the real-time log configuration to delete.")
	cloudfront_deleteRealtimeLogConfig2020_05_31Cmd.Flags().String("name", "", "The name of the real-time log configuration to delete.")
	cloudfrontCmd.AddCommand(cloudfront_deleteRealtimeLogConfig2020_05_31Cmd)
}
