package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoArchivedMediaCmd = &cobra.Command{
	Use:   "kinesis-video-archived-media",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoArchivedMediaCmd).Standalone()

	rootCmd.AddCommand(kinesisVideoArchivedMediaCmd)
}
