package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_startApplicationCmd = &cobra.Command{
	Use:   "start-application",
	Short: "Starts a specified application and initializes initial capacity if configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_startApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrServerless_startApplicationCmd).Standalone()

		emrServerless_startApplicationCmd.Flags().String("application-id", "", "The ID of the application to start.")
		emrServerless_startApplicationCmd.MarkFlagRequired("application-id")
	})
	emrServerlessCmd.AddCommand(emrServerless_startApplicationCmd)
}
