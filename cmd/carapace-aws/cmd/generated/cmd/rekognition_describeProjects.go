package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_describeProjectsCmd = &cobra.Command{
	Use:   "describe-projects",
	Short: "Gets information about your Rekognition projects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_describeProjectsCmd).Standalone()

	rekognition_describeProjectsCmd.Flags().String("features", "", "Specifies the type of customization to filter projects by.")
	rekognition_describeProjectsCmd.Flags().String("max-results", "", "The maximum number of results to return per paginated call.")
	rekognition_describeProjectsCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more results to retrieve), Rekognition returns a pagination token in the response.")
	rekognition_describeProjectsCmd.Flags().String("project-names", "", "A list of the projects that you want Rekognition to describe.")
	rekognitionCmd.AddCommand(rekognition_describeProjectsCmd)
}
