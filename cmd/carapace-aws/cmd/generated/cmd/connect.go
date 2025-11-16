package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "- [Amazon Connect actions](https://docs.aws.amazon.com/connect/latest/APIReference/API_Operations_Amazon_Connect_Service.html)\n- [Amazon Connect data types](https://docs.aws.amazon.com/connect/latest/APIReference/API_Types_Amazon_Connect_Service.html)\n\nAmazon Connect is a cloud-based contact center solution that you use to set up and manage a customer contact center and provide reliable customer engagement at any scale.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectCmd).Standalone()

	rootCmd.AddCommand(connectCmd)
}
