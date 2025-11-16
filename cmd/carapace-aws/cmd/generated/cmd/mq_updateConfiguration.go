package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_updateConfigurationCmd = &cobra.Command{
	Use:   "update-configuration",
	Short: "Updates the specified configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_updateConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_updateConfigurationCmd).Standalone()

		mq_updateConfigurationCmd.Flags().String("configuration-id", "", "The unique ID that Amazon MQ generates for the configuration.")
		mq_updateConfigurationCmd.Flags().String("data", "", "Amazon MQ for Active MQ: The base64-encoded XML configuration.")
		mq_updateConfigurationCmd.Flags().String("description", "", "The description of the configuration.")
		mq_updateConfigurationCmd.MarkFlagRequired("configuration-id")
		mq_updateConfigurationCmd.MarkFlagRequired("data")
	})
	mqCmd.AddCommand(mq_updateConfigurationCmd)
}
