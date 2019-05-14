package template

// Template file structure of module
var Template = map[string]map[string]string{
	"__init__.py": map[string]string{
		"content": RootInit,
	},
	"__manifest__.py": map[string]string{
		"content": Manifest,
	},
	"controllers": map[string]string{
		"__init__.py":    CInit,
		"controllers.py": Controllers,
	},
	"models": map[string]string{
		"__init__.py": MInit,
		"models.py":   Model,
	},
	"views": map[string]string{
		"templates.xml": ViewTemplates,
		"views.xml":     ViewViews,
	},
	"demo": map[string]string{
		"demo.xml": Demo,
	},
	"security": map[string]string{
		"ir.model.access.csv": Security,
	},
}
