package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_getCompositionCmd = &cobra.Command{
	Use:   "get-composition",
	Short: "Get information about the specified Composition resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_getCompositionCmd).Standalone()

	ivsRealtime_getCompositionCmd.Flags().String("arn", "", "ARN of the Composition resource.")
	ivsRealtime_getCompositionCmd.MarkFlagRequired("arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_getCompositionCmd)
}
