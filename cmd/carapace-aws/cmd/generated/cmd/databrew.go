package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrewCmd = &cobra.Command{
	Use:   "databrew",
	Short: "Glue DataBrew is a visual, cloud-scale data-preparation service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrewCmd).Standalone()

	rootCmd.AddCommand(databrewCmd)
}
