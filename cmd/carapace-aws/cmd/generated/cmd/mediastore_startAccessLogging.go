package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_startAccessLoggingCmd = &cobra.Command{
	Use:   "start-access-logging",
	Short: "Starts access logging on the specified container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_startAccessLoggingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_startAccessLoggingCmd).Standalone()

		mediastore_startAccessLoggingCmd.Flags().String("container-name", "", "The name of the container that you want to start access logging on.")
		mediastore_startAccessLoggingCmd.MarkFlagRequired("container-name")
	})
	mediastoreCmd.AddCommand(mediastore_startAccessLoggingCmd)
}
