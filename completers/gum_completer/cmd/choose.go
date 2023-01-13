package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
)

var chooseCmd = &cobra.Command{
	Use:   "choose",
	Short: "Choose an option from a list of choices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chooseCmd).Standalone()
	chooseCmd.Flags().String("cursor", "", "Prefix to show on item that corresponds to the cursor position")
	chooseCmd.Flags().String("cursor-prefix", "", "Prefix to show on the cursor item (hidden if limit is 1)")
	chooseCmd.Flags().String("cursor.align", "", "Text Alignment")
	chooseCmd.Flags().String("cursor.background", "", "Background Color")
	chooseCmd.Flags().Bool("cursor.bold", false, "Bold text")
	chooseCmd.Flags().String("cursor.border", "", "Border Style")
	chooseCmd.Flags().String("cursor.border-background", "", "Border Background Color")
	chooseCmd.Flags().String("cursor.border-foreground", "", "Border Foreground Color")
	chooseCmd.Flags().Bool("cursor.faint", false, "Faint text")
	chooseCmd.Flags().String("cursor.foreground", "", "Foreground Color")
	chooseCmd.Flags().String("cursor.height", "", "Text height")
	chooseCmd.Flags().Bool("cursor.italic", false, "Italicize text")
	chooseCmd.Flags().String("cursor.margin", "", "Text margin")
	chooseCmd.Flags().String("cursor.padding", "", "Text padding")
	chooseCmd.Flags().Bool("cursor.strikethrough", false, "Strikethrough text")
	chooseCmd.Flags().Bool("cursor.underline", false, "Underline text")
	chooseCmd.Flags().String("cursor.width", "", "Text width")
	chooseCmd.Flags().String("height", "", "Height of the list")
	chooseCmd.Flags().String("item.align", "", "Text Alignment")
	chooseCmd.Flags().String("item.background", "", "Background Color")
	chooseCmd.Flags().Bool("item.bold", false, "Bold text")
	chooseCmd.Flags().String("item.border", "", "Border Style")
	chooseCmd.Flags().String("item.border-background", "", "Border Background Color")
	chooseCmd.Flags().String("item.border-foreground", "", "Border Foreground Color")
	chooseCmd.Flags().Bool("item.faint", false, "Faint text")
	chooseCmd.Flags().String("item.foreground", "", "Foreground Color")
	chooseCmd.Flags().String("item.height", "", "Text height")
	chooseCmd.Flags().Bool("item.italic", false, "Italicize text")
	chooseCmd.Flags().String("item.margin", "", "Text margin")
	chooseCmd.Flags().String("item.padding", "", "Text padding")
	chooseCmd.Flags().Bool("item.strikethrough", false, "Strikethrough text")
	chooseCmd.Flags().Bool("item.underline", false, "Underline text")
	chooseCmd.Flags().String("item.width", "", "Text width")
	chooseCmd.Flags().String("limit", "", "Maximum number of options to pick")
	chooseCmd.Flags().Bool("no-limit", false, "Pick unlimited number of options (ignores limit)")
	chooseCmd.Flags().StringSlice("selected", []string{}, "Options that should start as selected")
	chooseCmd.Flags().String("selected-prefix", "", "Prefix to show on selected items (hidden if limit is 1)")
	chooseCmd.Flags().String("selected.align", "", "Text Alignment")
	chooseCmd.Flags().String("selected.background", "", "Background Color")
	chooseCmd.Flags().Bool("selected.bold", false, "Bold text")
	chooseCmd.Flags().String("selected.border", "", "Border Style")
	chooseCmd.Flags().String("selected.border-background", "", "Border Background Color")
	chooseCmd.Flags().String("selected.border-foreground", "", "Border Foreground Color")
	chooseCmd.Flags().Bool("selected.faint", false, "Faint text")
	chooseCmd.Flags().String("selected.foreground", "", "Foreground Color")
	chooseCmd.Flags().String("selected.height", "", "Text height")
	chooseCmd.Flags().Bool("selected.italic", false, "Italicize text")
	chooseCmd.Flags().String("selected.margin", "", "Text margin")
	chooseCmd.Flags().String("selected.padding", "", "Text padding")
	chooseCmd.Flags().Bool("selected.strikethrough", false, "Strikethrough text")
	chooseCmd.Flags().Bool("selected.underline", false, "Underline text")
	chooseCmd.Flags().String("selected.width", "", "Text width")
	chooseCmd.Flags().String("unselected-prefix", "", "Prefix to show on unselected items (hidden if limit is 1)")
	rootCmd.AddCommand(chooseCmd)

	carapace.Gen(chooseCmd).FlagCompletion(carapace.ActionMap{
		"cursor.align":             gum.ActionAlignments(),
		"cursor.background":        gum.ActionColors(),
		"cursor.border":            gum.ActionBorders(),
		"cursor.border-background": gum.ActionColors(),
		"cursor.border-foreground": gum.ActionColors(),
		"cursor.foreground":        gum.ActionColors(),
		"item.align":               gum.ActionAlignments(),
		"item.background":          gum.ActionColors(),
		"item.border":              gum.ActionBorders(),
		"item.border-background":   gum.ActionColors(),
		"item.border-foreground":   gum.ActionColors(),
		"item.foreground":          gum.ActionColors(),
		"selected": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionValues(c.Args...).UniqueList(",")
		}),
		"selected.align":             gum.ActionAlignments(),
		"selected.background":        gum.ActionColors(),
		"selected.border":            gum.ActionBorders(),
		"selected.border-background": gum.ActionColors(),
		"selected.border-foreground": gum.ActionColors(),
		"selected.foreground":        gum.ActionColors(),
	})
}
