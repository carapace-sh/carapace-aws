package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateRealtimeLogConfig2020_05_31Cmd = &cobra.Command{
	Use:   "update-realtime-log-config2020_05_31",
	Short: "Updates a real-time log configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateRealtimeLogConfig2020_05_31Cmd).Standalone()

	cloudfront_updateRealtimeLogConfig2020_05_31Cmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) for this real-time log configuration.")
	cloudfront_updateRealtimeLogConfig2020_05_31Cmd.Flags().String("end-points", "", "Contains information about the Amazon Kinesis data stream where you are sending real-time log data.")
	cloudfront_updateRealtimeLogConfig2020_05_31Cmd.Flags().String("fields", "", "A list of fields to include in each real-time log record.")
	cloudfront_updateRealtimeLogConfig2020_05_31Cmd.Flags().String("name", "", "The name for this real-time log configuration.")
	cloudfront_updateRealtimeLogConfig2020_05_31Cmd.Flags().String("sampling-rate", "", "The sampling rate for this real-time log configuration.")
	cloudfrontCmd.AddCommand(cloudfront_updateRealtimeLogConfig2020_05_31Cmd)
}
