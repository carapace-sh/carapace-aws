package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_stopApplicationCmd = &cobra.Command{
	Use:   "stop-application",
	Short: "Stops a specified application and releases initial capacity if configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_stopApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrServerless_stopApplicationCmd).Standalone()

		emrServerless_stopApplicationCmd.Flags().String("application-id", "", "The ID of the application to stop.")
		emrServerless_stopApplicationCmd.MarkFlagRequired("application-id")
	})
	emrServerlessCmd.AddCommand(emrServerless_stopApplicationCmd)
}
