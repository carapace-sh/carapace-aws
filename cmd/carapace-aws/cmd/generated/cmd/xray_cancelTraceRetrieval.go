package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_cancelTraceRetrievalCmd = &cobra.Command{
	Use:   "cancel-trace-retrieval",
	Short: "Cancels an ongoing trace retrieval job initiated by `StartTraceRetrieval` using the provided `RetrievalToken`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_cancelTraceRetrievalCmd).Standalone()

	xray_cancelTraceRetrievalCmd.Flags().String("retrieval-token", "", "Retrieval token.")
	xray_cancelTraceRetrievalCmd.MarkFlagRequired("retrieval-token")
	xrayCmd.AddCommand(xray_cancelTraceRetrievalCmd)
}
