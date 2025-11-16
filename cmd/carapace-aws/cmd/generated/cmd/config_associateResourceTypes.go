package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_associateResourceTypesCmd = &cobra.Command{
	Use:   "associate-resource-types",
	Short: "Adds all resource types specified in the `ResourceTypes` list to the [RecordingGroup](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html) of specified configuration recorder and includes those resource types when recording.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_associateResourceTypesCmd).Standalone()

	config_associateResourceTypesCmd.Flags().String("configuration-recorder-arn", "", "The Amazon Resource Name (ARN) of the specified configuration recorder.")
	config_associateResourceTypesCmd.Flags().String("resource-types", "", "The list of resource types you want to add to the recording group of the specified configuration recorder.")
	config_associateResourceTypesCmd.MarkFlagRequired("configuration-recorder-arn")
	config_associateResourceTypesCmd.MarkFlagRequired("resource-types")
	configCmd.AddCommand(config_associateResourceTypesCmd)
}
