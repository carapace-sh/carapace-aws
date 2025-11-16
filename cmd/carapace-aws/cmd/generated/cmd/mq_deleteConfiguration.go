package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_deleteConfigurationCmd = &cobra.Command{
	Use:   "delete-configuration",
	Short: "Deletes the specified configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_deleteConfigurationCmd).Standalone()

	mq_deleteConfigurationCmd.Flags().String("configuration-id", "", "The unique ID that Amazon MQ generates for the configuration.")
	mq_deleteConfigurationCmd.MarkFlagRequired("configuration-id")
	mqCmd.AddCommand(mq_deleteConfigurationCmd)
}
