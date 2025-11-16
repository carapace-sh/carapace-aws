package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_describeConfigurationRevisionCmd = &cobra.Command{
	Use:   "describe-configuration-revision",
	Short: "Returns the specified configuration revision for the specified configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_describeConfigurationRevisionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_describeConfigurationRevisionCmd).Standalone()

		mq_describeConfigurationRevisionCmd.Flags().String("configuration-id", "", "The unique ID that Amazon MQ generates for the configuration.")
		mq_describeConfigurationRevisionCmd.Flags().String("configuration-revision", "", "The revision of the configuration.")
		mq_describeConfigurationRevisionCmd.MarkFlagRequired("configuration-id")
		mq_describeConfigurationRevisionCmd.MarkFlagRequired("configuration-revision")
	})
	mqCmd.AddCommand(mq_describeConfigurationRevisionCmd)
}
