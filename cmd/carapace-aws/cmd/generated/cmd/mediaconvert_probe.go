package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_probeCmd = &cobra.Command{
	Use:   "probe",
	Short: "Use Probe to obtain detailed information about your input media files.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_probeCmd).Standalone()

	mediaconvert_probeCmd.Flags().String("input-files", "", "Specify a media file to probe.")
	mediaconvertCmd.AddCommand(mediaconvert_probeCmd)
}
