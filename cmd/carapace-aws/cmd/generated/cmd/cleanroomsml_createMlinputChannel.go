package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_createMlinputChannelCmd = &cobra.Command{
	Use:   "create-mlinput-channel",
	Short: "Provides the information to create an ML input channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_createMlinputChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_createMlinputChannelCmd).Standalone()

		cleanroomsml_createMlinputChannelCmd.Flags().String("configured-model-algorithm-associations", "", "The associated configured model algorithms that are necessary to create this ML input channel.")
		cleanroomsml_createMlinputChannelCmd.Flags().String("description", "", "The description of the ML input channel.")
		cleanroomsml_createMlinputChannelCmd.Flags().String("input-channel", "", "The input data that is used to create this ML input channel.")
		cleanroomsml_createMlinputChannelCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key that is used to access the input channel.")
		cleanroomsml_createMlinputChannelCmd.Flags().String("membership-identifier", "", "The membership ID of the member that is creating the ML input channel.")
		cleanroomsml_createMlinputChannelCmd.Flags().String("name", "", "The name of the ML input channel.")
		cleanroomsml_createMlinputChannelCmd.Flags().String("retention-in-days", "", "The number of days that the data in the ML input channel is retained.")
		cleanroomsml_createMlinputChannelCmd.Flags().String("tags", "", "The optional metadata that you apply to the resource to help you categorize and organize them.")
		cleanroomsml_createMlinputChannelCmd.MarkFlagRequired("configured-model-algorithm-associations")
		cleanroomsml_createMlinputChannelCmd.MarkFlagRequired("input-channel")
		cleanroomsml_createMlinputChannelCmd.MarkFlagRequired("membership-identifier")
		cleanroomsml_createMlinputChannelCmd.MarkFlagRequired("name")
		cleanroomsml_createMlinputChannelCmd.MarkFlagRequired("retention-in-days")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_createMlinputChannelCmd)
}
