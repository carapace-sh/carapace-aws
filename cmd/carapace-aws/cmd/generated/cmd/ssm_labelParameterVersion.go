package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_labelParameterVersionCmd = &cobra.Command{
	Use:   "label-parameter-version",
	Short: "A parameter label is a user-defined alias to help you manage different versions of a parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_labelParameterVersionCmd).Standalone()

	ssm_labelParameterVersionCmd.Flags().String("labels", "", "One or more labels to attach to the specified parameter version.")
	ssm_labelParameterVersionCmd.Flags().String("name", "", "The parameter name on which you want to attach one or more labels.")
	ssm_labelParameterVersionCmd.Flags().String("parameter-version", "", "The specific version of the parameter on which you want to attach one or more labels.")
	ssm_labelParameterVersionCmd.MarkFlagRequired("labels")
	ssm_labelParameterVersionCmd.MarkFlagRequired("name")
	ssmCmd.AddCommand(ssm_labelParameterVersionCmd)
}
