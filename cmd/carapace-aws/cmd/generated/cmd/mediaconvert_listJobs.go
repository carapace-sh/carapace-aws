package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_listJobsCmd = &cobra.Command{
	Use:   "list-jobs",
	Short: "Retrieve a JSON array of up to twenty of your most recently created jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_listJobsCmd).Standalone()

	mediaconvert_listJobsCmd.Flags().String("max-results", "", "Optional.")
	mediaconvert_listJobsCmd.Flags().String("next-token", "", "Optional.")
	mediaconvert_listJobsCmd.Flags().String("order", "", "Optional.")
	mediaconvert_listJobsCmd.Flags().String("queue", "", "Optional.")
	mediaconvert_listJobsCmd.Flags().String("status", "", "Optional.")
	mediaconvertCmd.AddCommand(mediaconvert_listJobsCmd)
}
