package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_createConfigurationCmd = &cobra.Command{
	Use:   "create-configuration",
	Short: "Creates a new configuration for the specified configuration name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_createConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_createConfigurationCmd).Standalone()

		mq_createConfigurationCmd.Flags().String("authentication-strategy", "", "Optional.")
		mq_createConfigurationCmd.Flags().String("engine-type", "", "Required.")
		mq_createConfigurationCmd.Flags().String("engine-version", "", "The broker engine version.")
		mq_createConfigurationCmd.Flags().String("name", "", "Required.")
		mq_createConfigurationCmd.Flags().String("tags", "", "Create tags when creating the configuration.")
		mq_createConfigurationCmd.MarkFlagRequired("engine-type")
		mq_createConfigurationCmd.MarkFlagRequired("name")
	})
	mqCmd.AddCommand(mq_createConfigurationCmd)
}
