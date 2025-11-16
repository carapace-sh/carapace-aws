package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtimeCmd = &cobra.Command{
	Use:   "ivs-realtime",
	Short: "The Amazon Interactive Video Service (IVS) real-time API is REST compatible, using a standard HTTP API and an AWS EventBridge event stream for responses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtimeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtimeCmd).Standalone()

	})
	rootCmd.AddCommand(ivsRealtimeCmd)
}
