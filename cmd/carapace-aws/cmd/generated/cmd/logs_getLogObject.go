package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getLogObjectCmd = &cobra.Command{
	Use:   "get-log-object",
	Short: "Retrieves a large logging object (LLO) and streams it back.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getLogObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_getLogObjectCmd).Standalone()

		logs_getLogObjectCmd.Flags().String("log-object-pointer", "", "A pointer to the specific log object to retrieve.")
		logs_getLogObjectCmd.Flags().String("unmask", "", "A boolean flag that indicates whether to unmask sensitive log data.")
		logs_getLogObjectCmd.MarkFlagRequired("log-object-pointer")
	})
	logsCmd.AddCommand(logs_getLogObjectCmd)
}
