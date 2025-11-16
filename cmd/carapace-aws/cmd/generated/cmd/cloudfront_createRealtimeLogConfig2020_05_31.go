package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createRealtimeLogConfig2020_05_31Cmd = &cobra.Command{
	Use:   "create-realtime-log-config2020_05_31",
	Short: "Creates a real-time log configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createRealtimeLogConfig2020_05_31Cmd).Standalone()

	cloudfront_createRealtimeLogConfig2020_05_31Cmd.Flags().String("end-points", "", "Contains information about the Amazon Kinesis data stream where you are sending real-time log data.")
	cloudfront_createRealtimeLogConfig2020_05_31Cmd.Flags().String("fields", "", "A list of fields to include in each real-time log record.")
	cloudfront_createRealtimeLogConfig2020_05_31Cmd.Flags().String("name", "", "A unique name to identify this real-time log configuration.")
	cloudfront_createRealtimeLogConfig2020_05_31Cmd.Flags().String("sampling-rate", "", "The sampling rate for this real-time log configuration.")
	cloudfront_createRealtimeLogConfig2020_05_31Cmd.MarkFlagRequired("end-points")
	cloudfront_createRealtimeLogConfig2020_05_31Cmd.MarkFlagRequired("fields")
	cloudfront_createRealtimeLogConfig2020_05_31Cmd.MarkFlagRequired("name")
	cloudfront_createRealtimeLogConfig2020_05_31Cmd.MarkFlagRequired("sampling-rate")
	cloudfrontCmd.AddCommand(cloudfront_createRealtimeLogConfig2020_05_31Cmd)
}
