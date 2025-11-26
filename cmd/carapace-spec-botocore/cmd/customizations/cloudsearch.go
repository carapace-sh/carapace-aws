package customizations

import "github.com/carapace-sh/carapace-spec/pkg/command"

func init() {
	customizations["cloudsearch.define-expression"] = func(cmd *command.Command) error {
		cmd.AddFlag(command.Flag{
			Longhand:    "name",
			Description: "Names must begin with a letter and can contain the following characters: a-z (lowercase), 0-9, and _ (underscore).",
			Value:       true,
			Required:    true,
		})
		return nil
	}

	customizations["cloudsearch.define-index-field"] = func(cmd *command.Command) error {
		delete(cmd.Flags, "--index-field=!")

		cmd.AddFlag(command.Flag{
			Longhand:    "analysis-scheme",
			Description: "The name of an analysis scheme for a text field.",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "facet-enabled",
			Description: "Whether facet information can be returned for the field.",
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "sort-enabled",
			Description: "Whether the field can be used to sort the search results.",
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "return-enabled",
			Description: "Whether the contents of the field can be returned in the search results.",
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "default-value",
			Description: "A value to use for the field if the field isn’t specified for a document.",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "source-field",
			Description: "The name of the source field to map to the field.",
			Value:       true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "highlight-enabled",
			Description: "Whether highlights can be returned for the field.",
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "name",
			Description: "A string that represents the name of an index field",
			Value:       true,
			Required:    true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "type",
			Description: "The type of field.",
			Value:       true,
			Required:    true,
		})
		cmd.AddFlag(command.Flag{
			Longhand:    "search-enabled",
			Description: "Whether highlights can be returned for the field.",
		})

		// TODO documentation
		return nil
	}
}
