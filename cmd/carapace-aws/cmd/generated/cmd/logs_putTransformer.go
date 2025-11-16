package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putTransformerCmd = &cobra.Command{
	Use:   "put-transformer",
	Short: "Creates or updates a *log transformer* for a single log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putTransformerCmd).Standalone()

	logs_putTransformerCmd.Flags().String("log-group-identifier", "", "Specify either the name or ARN of the log group to create the transformer for.")
	logs_putTransformerCmd.Flags().String("transformer-config", "", "This structure contains the configuration of this log transformer.")
	logs_putTransformerCmd.MarkFlagRequired("log-group-identifier")
	logs_putTransformerCmd.MarkFlagRequired("transformer-config")
	logsCmd.AddCommand(logs_putTransformerCmd)
}
