package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeContainerGroupDefinitionCmd = &cobra.Command{
	Use:   "describe-container-group-definition",
	Short: "**This API works with the following fleet types:** Container\n\nRetrieves the properties of a container group definition, including all container definitions in the group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeContainerGroupDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_describeContainerGroupDefinitionCmd).Standalone()

		gamelift_describeContainerGroupDefinitionCmd.Flags().String("name", "", "The unique identifier for the container group definition to retrieve properties for.")
		gamelift_describeContainerGroupDefinitionCmd.Flags().String("version-number", "", "The specific version to retrieve.")
		gamelift_describeContainerGroupDefinitionCmd.MarkFlagRequired("name")
	})
	gameliftCmd.AddCommand(gamelift_describeContainerGroupDefinitionCmd)
}
