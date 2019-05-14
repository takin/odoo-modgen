package cmd

import (
	"fmt"
	"os"
	"regexp"

	"github.com/takin/odoo-modgen/helper"
	t "github.com/takin/odoo-modgen/template"
)

// OdooModule module model
type OdooModule struct {
	Name        string
	Path        string
	OdooVersion int
}

func convertTemplateContent(moduleName, template string) string {
	r, err := regexp.Compile("(\\{MOD_NAME\\})(\\{MOD_NAME_LOWER\\})")
	if err != nil {
		fmt.Errorf("ERROR: can not compile template converter\nTemplate will be using default content")
		return template
	}
	replacedTemplate := r.ReplaceAllStringFunc(template, func(res string) string {
		fmt.Println(res)
		return res
	})
	return replacedTemplate
}

func (o *OdooModule) Generate() {
	prepare(o)

	projectpath := o.Path + "/" + o.Name

	switch o.OdooVersion {
	case 12:

		for k, v := range t.Template {
			match, err := regexp.MatchString("\\.py$", k)
			if err != nil {
				fmt.Errorf("failed to compare entry format of: %s\n", k)
				os.Exit(1)
			}
			// Jika match, maka ini adalah file, bukan direktori
			if match {
				for _, iv := range v {
					// buat file dengan nama file sesuai dengan nilai k saat ini
					// misal __manifest__.py
					// dengan content berdasarkan nilai dari map["content"]
					// dari __manifest__.py:map[string]string{"content":<file content>}
					helper.CreateFile(projectpath+"/"+k, iv)
				}
			} else {
				newDir := projectpath + "/" + k
				// buat folder untuk masing-masing part module
				// controllers, models, etc
				helper.CreateDir(newDir, 0755, false)
				// setelah directory nya dibuat, maka selanjutnya buat file
				// content dari direktori yang bersangkutan
				// sesuai dengan isi dari Map folder yang bersangkutan
				// "controllers":map[string]string{
				// 		"__init__.py":<init content>
				// }
				for ik, iv := range v {
					helper.CreateFile(newDir+"/"+ik, iv)
				}
			}
		}

	}
}

func prepare(o *OdooModule) {

	projectDir := o.Path + "/" + o.Name

	if _, err := os.Stat(o.Path); os.IsNotExist(err) {
		helper.Confirm("Specified path does not exist, do you want to create it? [yes|no]")
		helper.CreateDir(o.Path, 0755, true)
	}

	// check if the module name as a directory existed
	if _, err := os.Stat(projectDir); os.IsExist(err) {
		helper.Confirm(fmt.Sprintf("%s directory existed, do you want to replace?[yes|no]", o.Name))
	}

	helper.CreateDir(projectDir, 0755, false)
}
