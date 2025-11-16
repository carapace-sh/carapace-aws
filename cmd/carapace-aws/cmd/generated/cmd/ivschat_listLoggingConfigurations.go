package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivschat_listLoggingConfigurationsCmd = &cobra.Command{
	Use:   "list-logging-configurations",
	Short: "Gets summary information about all your logging configurations in the AWS region where the API request is processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivschat_listLoggingConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivschat_listLoggingConfigurationsCmd).Standalone()

		ivschat_listLoggingConfigurationsCmd.Flags().String("max-results", "", "Maximum number of logging configurations to return.")
		ivschat_listLoggingConfigurationsCmd.Flags().String("next-token", "", "The first logging configurations to retrieve.")
	})
	ivschatCmd.AddCommand(ivschat_listLoggingConfigurationsCmd)
}
