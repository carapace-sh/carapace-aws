package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_disassociateResourceTypesCmd = &cobra.Command{
	Use:   "disassociate-resource-types",
	Short: "Removes all resource types specified in the `ResourceTypes` list from the [RecordingGroup](https://docs.aws.amazon.com/config/latest/APIReference/API_RecordingGroup.html) of configuration recorder and excludes these resource types when recording.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_disassociateResourceTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_disassociateResourceTypesCmd).Standalone()

		config_disassociateResourceTypesCmd.Flags().String("configuration-recorder-arn", "", "The Amazon Resource Name (ARN) of the specified configuration recorder.")
		config_disassociateResourceTypesCmd.Flags().String("resource-types", "", "The list of resource types you want to remove from the recording group of the specified configuration recorder.")
		config_disassociateResourceTypesCmd.MarkFlagRequired("configuration-recorder-arn")
		config_disassociateResourceTypesCmd.MarkFlagRequired("resource-types")
	})
	configCmd.AddCommand(config_disassociateResourceTypesCmd)
}
