package {{.Model}}

{{$ilen := len .Imports}}
{{if gt $ilen 0}}
import (
	{{range .Imports}}"{{.}}"{{end}}
)
{{end}}
{{range .Tables}}
type {{Mapper .Name}} struct {
{{$table := .}}
{{range .ColumnsSeq}}{{$col := $table.GetColumn .}}	{{Mapper $col.Name}}  {{Type $col}}  {{Tag $table $col}}
{{end}}
}

func (m {{Mapper .Name}}) TableName() string {
	return "ty_{{$table.Name}}"
}

func (m *{{Mapper .Name}}) AfterFind(){
}

func Add{{Mapper .Name}}(model *{{Mapper .Name}}) error {
	return AddModel(model)
}

func Del{{Mapper .Name}}ById(id int64, safe ...bool) error {
	model := new({{Mapper .Name}})
	return DelModel(model, Engine.ID(id), safe...)
}

func Del{{Mapper .Name}}(session *xorm.Session, safe ...bool) error {
	model := new({{Mapper .Name}})
	return DelModel(model, session, safe...)
}

func Update{{Mapper .Name}}ById(id int64, model *{{Mapper .Name}}) error {
	return UpdateModel(model, Engine.ID(id))
}

func Update{{Mapper .Name}}(model *{{Mapper .Name}}, session *xorm.Session) error {
	return UpdateModel(model, session)
}

func Get{{Mapper .Name}}ByID(id int64) (*{{Mapper .Name}}, error) {
	model := new({{Mapper .Name}})

	return model, GetModel(model, Engine.ID(id))
}

func Get{{Mapper .Name}}(session *xorm.Session) (*{{Mapper .Name}}, error) {
	model := new({{Mapper .Name}})

	return model, GetModel(model, session)
}

func Get{{Mapper .Name}}List(session *xorm.Session, limit ...int) (models []*{{Mapper .Name}}, err error) {
	if len(limit) > 0 {
		models = make([]*{{Mapper .Name}}, 0, limit[0])

		err = session.Limit(limit[0]).Find(&models)

	} else {
		models = make([]*{{Mapper .Name}}, 0)

		err = session.Find(&models)

	}

	if err != nil {
		return nil, err
	}

	for i := range models {
		models[i].AfterFind()
	}

	return models, nil
}

func Get{{Mapper .Name}}ListForPage(session *xorm.Session, page int, pageSize int) ([]*{{Mapper .Name}}, error) {
	models := make([]*{{Mapper .Name}}, 0, pageSize)

	err := session.Limit(pageSize, (page-1)*pageSize).Find(&models)

	if err != nil {
		return nil, err
	}

	for i := range models {
		models[i].AfterFind()
	}

	return models, nil
}

{{end}}