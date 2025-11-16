package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_describeConfigurationCmd = &cobra.Command{
	Use:   "describe-configuration",
	Short: "Returns information about the specified configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_describeConfigurationCmd).Standalone()

	mq_describeConfigurationCmd.Flags().String("configuration-id", "", "The unique ID that Amazon MQ generates for the configuration.")
	mq_describeConfigurationCmd.MarkFlagRequired("configuration-id")
	mqCmd.AddCommand(mq_describeConfigurationCmd)
}
