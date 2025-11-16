package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getMasterAccountCmd = &cobra.Command{
	Use:   "get-master-account",
	Short: "This method is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getMasterAccountCmd).Standalone()

	securityhubCmd.AddCommand(securityhub_getMasterAccountCmd)
}
