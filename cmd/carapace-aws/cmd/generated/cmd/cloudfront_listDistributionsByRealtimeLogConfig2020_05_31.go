package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDistributionsByRealtimeLogConfig2020_05_31Cmd = &cobra.Command{
	Use:   "list-distributions-by-realtime-log-config2020_05_31",
	Short: "Gets a list of distributions that have a cache behavior that's associated with the specified real-time log configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDistributionsByRealtimeLogConfig2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listDistributionsByRealtimeLogConfig2020_05_31Cmd).Standalone()

		cloudfront_listDistributionsByRealtimeLogConfig2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of distributions.")
		cloudfront_listDistributionsByRealtimeLogConfig2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of distributions that you want in the response.")
		cloudfront_listDistributionsByRealtimeLogConfig2020_05_31Cmd.Flags().String("realtime-log-config-arn", "", "The Amazon Resource Name (ARN) of the real-time log configuration whose associated distributions you want to list.")
		cloudfront_listDistributionsByRealtimeLogConfig2020_05_31Cmd.Flags().String("realtime-log-config-name", "", "The name of the real-time log configuration whose associated distributions you want to list.")
	})
	cloudfrontCmd.AddCommand(cloudfront_listDistributionsByRealtimeLogConfig2020_05_31Cmd)
}
