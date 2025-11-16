package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_generateAccessLogsCmd = &cobra.Command{
	Use:   "generate-access-logs",
	Short: "Returns the website access logs for a specific time range using a presigned URL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_generateAccessLogsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplify_generateAccessLogsCmd).Standalone()

		amplify_generateAccessLogsCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
		amplify_generateAccessLogsCmd.Flags().String("domain-name", "", "The name of the domain.")
		amplify_generateAccessLogsCmd.Flags().String("end-time", "", "The time at which the logs should end.")
		amplify_generateAccessLogsCmd.Flags().String("start-time", "", "The time at which the logs should start.")
		amplify_generateAccessLogsCmd.MarkFlagRequired("app-id")
		amplify_generateAccessLogsCmd.MarkFlagRequired("domain-name")
	})
	amplifyCmd.AddCommand(amplify_generateAccessLogsCmd)
}
