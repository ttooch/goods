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

func Get{{Mapper .Name}}ByID(id int64) (*{{Mapper .Name}}, error) {
    model := new({{Mapper .Name}})

    has, err := Engine.Id(id).Get(model)

    model.AfterFind()

    if err != nil {
        return model, err
    } else if !has {
        return model, ErrNotExist
    }

    return model, nil
}

func Get{{Mapper .Name}}(session *xorm.Session) (*{{Mapper .Name}}, error) {
	model := new({{Mapper .Name}})

	has, err := session.Get(model)

	model.AfterFind()

	if err != nil {
		return model, err
	} else if !has {
		return model, ErrNotExist
	}

	return model, nil
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

func Update{{Mapper .Name}}ById(id int64, model *{{Mapper .Name}}) (bool, error) {
	_, err := Engine.Id(id).Update(model)

	if err != nil {
		return false, err
	} else {
		return true, err
	}

}

func Update{{Mapper .Name}}(session *xorm.Session, model *{{Mapper .Name}}) (bool, error) {
	_, err := session.Update(model)

	if err != nil {
		return false, err
	} else {
		return true, err
	}

}

func Add{{Mapper .Name}}(model *{{Mapper .Name}}) error {
	effect, err := Engine.InsertOne(model)

	if err != nil {
		return err
	} else if effect == 0 {
		return ErrInsert
	}
	return nil
}

func Del{{Mapper .Name}}ById(id int64, safe ...bool) (deleted bool, err error) {
	var model {{Mapper .Name}}

	if len(safe) > 0 && safe[0] == false {
		_, err = Engine.ID(id).Unscoped().Delete(&model)
	} else {
		_, err = Engine.ID(id).Delete(&model)
	}

	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func Del{{Mapper .Name}}(session xorm.Session, safe ...bool) (deleted bool, err error) {
	var model {{Mapper .Name}}

	if len(safe) > 0 && safe[0] == false {
		_, err = session.Unscoped().Delete(&model)
	} else {
		_, err = session.Delete(&model)
	}

	if err != nil {
		return false, err
	} else {
		return true, err
	}
}

{{end}}