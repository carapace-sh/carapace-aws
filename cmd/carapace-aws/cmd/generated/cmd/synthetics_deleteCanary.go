package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_deleteCanaryCmd = &cobra.Command{
	Use:   "delete-canary",
	Short: "Permanently deletes the specified canary.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_deleteCanaryCmd).Standalone()

	synthetics_deleteCanaryCmd.Flags().String("delete-lambda", "", "Specifies whether to also delete the Lambda functions and layers used by this canary.")
	synthetics_deleteCanaryCmd.Flags().String("name", "", "The name of the canary that you want to delete.")
	synthetics_deleteCanaryCmd.MarkFlagRequired("name")
	syntheticsCmd.AddCommand(synthetics_deleteCanaryCmd)
}
