package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_unlabelParameterVersionCmd = &cobra.Command{
	Use:   "unlabel-parameter-version",
	Short: "Remove a label or labels from a parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_unlabelParameterVersionCmd).Standalone()

	ssm_unlabelParameterVersionCmd.Flags().String("labels", "", "One or more labels to delete from the specified parameter version.")
	ssm_unlabelParameterVersionCmd.Flags().String("name", "", "The name of the parameter from which you want to delete one or more labels.")
	ssm_unlabelParameterVersionCmd.Flags().String("parameter-version", "", "The specific version of the parameter which you want to delete one or more labels from.")
	ssm_unlabelParameterVersionCmd.MarkFlagRequired("labels")
	ssm_unlabelParameterVersionCmd.MarkFlagRequired("name")
	ssm_unlabelParameterVersionCmd.MarkFlagRequired("parameter-version")
	ssmCmd.AddCommand(ssm_unlabelParameterVersionCmd)
}
