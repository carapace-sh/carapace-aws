package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_putFeedbackCmd = &cobra.Command{
	Use:   "put-feedback",
	Short: "Collects customer feedback about the specified insight.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_putFeedbackCmd).Standalone()

	devopsGuru_putFeedbackCmd.Flags().String("insight-feedback", "", "The feedback from customers is about the recommendations in this insight.")
	devopsGuruCmd.AddCommand(devopsGuru_putFeedbackCmd)
}
