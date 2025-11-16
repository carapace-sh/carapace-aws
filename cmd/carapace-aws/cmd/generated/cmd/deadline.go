package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadlineCmd = &cobra.Command{
	Use:   "deadline",
	Short: "The Amazon Web Services Deadline Cloud API provides infrastructure and centralized management for your projects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadlineCmd).Standalone()

	rootCmd.AddCommand(deadlineCmd)
}
