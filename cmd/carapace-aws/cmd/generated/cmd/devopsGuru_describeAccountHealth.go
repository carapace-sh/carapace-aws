package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_describeAccountHealthCmd = &cobra.Command{
	Use:   "describe-account-health",
	Short: "Returns the number of open reactive insights, the number of open proactive insights, and the number of metrics analyzed in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_describeAccountHealthCmd).Standalone()

	devopsGuruCmd.AddCommand(devopsGuru_describeAccountHealthCmd)
}
