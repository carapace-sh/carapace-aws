package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_cancelJobCmd = &cobra.Command{
	Use:   "cancel-job",
	Short: "Permanently cancel a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_cancelJobCmd).Standalone()

	mediaconvert_cancelJobCmd.Flags().String("id", "", "The Job ID of the job to be cancelled.")
	mediaconvert_cancelJobCmd.MarkFlagRequired("id")
	mediaconvertCmd.AddCommand(mediaconvert_cancelJobCmd)
}
