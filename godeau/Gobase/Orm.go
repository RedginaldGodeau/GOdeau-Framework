package Gobase

func (db *Database) Insert(entity *interface{}) {
	db.ORM.Create(entity)
}

func (db *Database) Select(entity *interface{}, search ...interface{}) {
	db.ORM.First(entity, search)
}

func (db *Database) Update(entity *interface{}, arg string, value interface{}) {
	db.ORM.Model(&entity).Update(arg, value)
}

func (db *Database) Updates(entity *interface{}, values map[string]interface{}) {
	db.ORM.Model(&entity).Updates(values)
}

func (db *Database) Deletes(entity *interface{}, id int) {
	db.ORM.Delete(entity, id)
}
