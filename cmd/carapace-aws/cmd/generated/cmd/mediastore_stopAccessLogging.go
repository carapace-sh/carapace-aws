package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_stopAccessLoggingCmd = &cobra.Command{
	Use:   "stop-access-logging",
	Short: "Stops access logging on the specified container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_stopAccessLoggingCmd).Standalone()

	mediastore_stopAccessLoggingCmd.Flags().String("container-name", "", "The name of the container that you want to stop access logging on.")
	mediastore_stopAccessLoggingCmd.MarkFlagRequired("container-name")
	mediastoreCmd.AddCommand(mediastore_stopAccessLoggingCmd)
}
