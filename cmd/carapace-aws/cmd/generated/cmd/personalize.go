package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalizeCmd = &cobra.Command{
	Use:   "personalize",
	Short: "Amazon Personalize is a machine learning service that makes it easy to add individualized recommendations to customers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalizeCmd).Standalone()

	rootCmd.AddCommand(personalizeCmd)
}
