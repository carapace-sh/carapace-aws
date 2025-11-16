package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stsCmd = &cobra.Command{
	Use:   "sts",
	Short: "Security Token Service\n\nSecurity Token Service (STS) enables you to request temporary, limited-privilege credentials for users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stsCmd).Standalone()

	rootCmd.AddCommand(stsCmd)
}
