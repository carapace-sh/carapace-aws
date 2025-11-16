package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_getJobCmd = &cobra.Command{
	Use:   "get-job",
	Short: "Retrieve the JSON for a specific transcoding job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_getJobCmd).Standalone()

	mediaconvert_getJobCmd.Flags().String("id", "", "the job ID of the job.")
	mediaconvert_getJobCmd.MarkFlagRequired("id")
	mediaconvertCmd.AddCommand(mediaconvert_getJobCmd)
}
