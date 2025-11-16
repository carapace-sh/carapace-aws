package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_listJobTemplatesCmd = &cobra.Command{
	Use:   "list-job-templates",
	Short: "Retrieve a JSON array of up to twenty of your job templates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_listJobTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_listJobTemplatesCmd).Standalone()

		mediaconvert_listJobTemplatesCmd.Flags().String("category", "", "Optionally, specify a job template category to limit responses to only job templates from that category.")
		mediaconvert_listJobTemplatesCmd.Flags().String("list-by", "", "Optional.")
		mediaconvert_listJobTemplatesCmd.Flags().String("max-results", "", "Optional.")
		mediaconvert_listJobTemplatesCmd.Flags().String("next-token", "", "Use this string, provided with the response to a previous request, to request the next batch of job templates.")
		mediaconvert_listJobTemplatesCmd.Flags().String("order", "", "Optional.")
	})
	mediaconvertCmd.AddCommand(mediaconvert_listJobTemplatesCmd)
}
