package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifybackendCmd = &cobra.Command{
	Use:   "amplifybackend",
	Short: "AWS Amplify Admin API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifybackendCmd).Standalone()

	rootCmd.AddCommand(amplifybackendCmd)
}
