package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_describeHubCmd = &cobra.Command{
	Use:   "describe-hub",
	Short: "Returns details about the Hub resource in your account, including the `HubArn` and the time when you enabled Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_describeHubCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_describeHubCmd).Standalone()

		securityhub_describeHubCmd.Flags().String("hub-arn", "", "The ARN of the Hub resource to retrieve.")
	})
	securityhubCmd.AddCommand(securityhub_describeHubCmd)
}
