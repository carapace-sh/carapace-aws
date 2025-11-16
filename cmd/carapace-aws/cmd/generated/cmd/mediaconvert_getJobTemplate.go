package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_getJobTemplateCmd = &cobra.Command{
	Use:   "get-job-template",
	Short: "Retrieve the JSON for a specific job template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_getJobTemplateCmd).Standalone()

	mediaconvert_getJobTemplateCmd.Flags().String("name", "", "The name of the job template.")
	mediaconvert_getJobTemplateCmd.MarkFlagRequired("name")
	mediaconvertCmd.AddCommand(mediaconvert_getJobTemplateCmd)
}
