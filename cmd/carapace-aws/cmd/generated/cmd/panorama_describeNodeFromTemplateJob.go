package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_describeNodeFromTemplateJobCmd = &cobra.Command{
	Use:   "describe-node-from-template-job",
	Short: "Returns information about a job to create a camera stream node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_describeNodeFromTemplateJobCmd).Standalone()

	panorama_describeNodeFromTemplateJobCmd.Flags().String("job-id", "", "The job's ID.")
	panorama_describeNodeFromTemplateJobCmd.MarkFlagRequired("job-id")
	panoramaCmd.AddCommand(panorama_describeNodeFromTemplateJobCmd)
}
