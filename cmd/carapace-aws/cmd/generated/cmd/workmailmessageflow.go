package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmailmessageflowCmd = &cobra.Command{
	Use:   "workmailmessageflow",
	Short: "The WorkMail Message Flow API provides access to email messages as they are being sent and received by a WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmailmessageflowCmd).Standalone()

	rootCmd.AddCommand(workmailmessageflowCmd)
}
