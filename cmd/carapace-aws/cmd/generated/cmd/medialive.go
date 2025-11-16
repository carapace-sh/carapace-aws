package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialiveCmd = &cobra.Command{
	Use:   "medialive",
	Short: "API for AWS Elemental MediaLive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialiveCmd).Standalone()

	rootCmd.AddCommand(medialiveCmd)
}
