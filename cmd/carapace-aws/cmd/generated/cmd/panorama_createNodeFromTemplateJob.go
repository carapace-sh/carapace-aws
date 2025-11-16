package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_createNodeFromTemplateJobCmd = &cobra.Command{
	Use:   "create-node-from-template-job",
	Short: "Creates a camera stream node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_createNodeFromTemplateJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_createNodeFromTemplateJobCmd).Standalone()

		panorama_createNodeFromTemplateJobCmd.Flags().String("job-tags", "", "Tags for the job.")
		panorama_createNodeFromTemplateJobCmd.Flags().String("node-description", "", "A description for the node.")
		panorama_createNodeFromTemplateJobCmd.Flags().String("node-name", "", "A name for the node.")
		panorama_createNodeFromTemplateJobCmd.Flags().String("output-package-name", "", "An output package name for the node.")
		panorama_createNodeFromTemplateJobCmd.Flags().String("output-package-version", "", "An output package version for the node.")
		panorama_createNodeFromTemplateJobCmd.Flags().String("template-parameters", "", "Template parameters for the node.")
		panorama_createNodeFromTemplateJobCmd.Flags().String("template-type", "", "The type of node.")
		panorama_createNodeFromTemplateJobCmd.MarkFlagRequired("node-name")
		panorama_createNodeFromTemplateJobCmd.MarkFlagRequired("output-package-name")
		panorama_createNodeFromTemplateJobCmd.MarkFlagRequired("output-package-version")
		panorama_createNodeFromTemplateJobCmd.MarkFlagRequired("template-parameters")
		panorama_createNodeFromTemplateJobCmd.MarkFlagRequired("template-type")
	})
	panoramaCmd.AddCommand(panorama_createNodeFromTemplateJobCmd)
}
